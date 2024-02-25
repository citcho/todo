package todo

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/hexisa_go_nal_todo/internal/common/auth"
)

type Todo struct {
	ulid      string
	userId    string
	title     string
	content   string
	completed int
}

func NewTodo(
	ulid string,
	userId string,
	title string,
	content string,
) (*Todo, error) {
	if utf8.RuneCountInString(title) > 99 {
		return &Todo{}, errors.New("Todoのタイトルは100文字未満で入力してください。")
	}
	if utf8.RuneCountInString(content) > 999 {
		return &Todo{}, errors.New("Todoの内容は1000文字未満で入力してください。")
	}

	t := &Todo{
		ulid:      ulid,
		userId:    userId,
		title:     title,
		content:   content,
		completed: 0,
	}

	return t, nil
}

func (t Todo) canCompleteTodo(id string) bool {
	return id == t.userId
}

func (t *Todo) Complete(ctx context.Context) error {
	id, ok := auth.GetUserID(ctx)
	if !ok {
		return errors.New("ユーザーIDを取得できませんでした。")
	}

	if !t.canCompleteTodo(id) {
		return errors.New("Todoの完了に失敗しました。")
	}

	t.completed = 1

	return nil
}

func ReConstructFromRepository(
	ulid string,
	userId string,
	title string,
	content string,
	completed int,
) *Todo {
	t := &Todo{
		ulid:      ulid,
		userId:    userId,
		title:     title,
		content:   content,
		completed: completed,
	}

	return t
}

func (t Todo) Ulid() string {
	return t.ulid
}

func (t Todo) UserId() string {
	return t.userId
}

func (t Todo) Title() string {
	return t.title
}

func (t Todo) Content() string {
	return t.content
}

func (t Todo) Completed() int {
	return t.completed
}
