package todo

import (
	"errors"
	"unicode/utf8"
)

type Todo struct {
	ulid    string
	userId  string
	title   string
	content string
	status  Status
}

//go:generate stringer -type Status
type Status int

const (
	_ Status = iota
	Yet
	Done
)

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
		ulid:    ulid,
		userId:  userId,
		title:   title,
		content: content,
		status:  Yet,
	}

	return t, nil
}

func ReConstructFromRepository(
	ulid string,
	userId string,
	title string,
	content string,
	status Status,
) *Todo {
	t := &Todo{
		ulid:    ulid,
		userId:  userId,
		title:   title,
		content: content,
		status:  status,
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

func (t Todo) Status() Status {
	return t.status
}
