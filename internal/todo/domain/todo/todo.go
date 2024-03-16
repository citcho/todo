package todo

import (
	"context"
	"errors"
	"unicode/utf8"

	"github.com/hexisa_go_nal_todo/internal/pkg/auth"
)

type Todo struct {
	id         string
	userId     string
	title      string
	content    string
	isComplete bool
}

func NewTodo(
	id string,
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
		id:         id,
		userId:     userId,
		title:      title,
		content:    content,
		isComplete: false,
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

	t.isComplete = true

	return nil
}

func ReConstructFromRepository(
	id string,
	userId string,
	title string,
	content string,
	isComplete bool,
) *Todo {
	t := &Todo{
		id:         id,
		userId:     userId,
		title:      title,
		content:    content,
		isComplete: isComplete,
	}

	return t
}

func (t Todo) Id() string {
	return t.id
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

func (t Todo) IsComplete() bool {
	return t.isComplete
}
