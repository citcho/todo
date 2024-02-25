package todo

func NewTodoFixture(ulid string, userUlid string, title string, content string, completed int) *Todo {
	return &Todo{
		ulid:      ulid,
		userId:    userUlid,
		title:     title,
		content:   content,
		completed: completed,
	}
}
