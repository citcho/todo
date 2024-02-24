package todo

func NewTodoFixture(ulid string, userUlid string, title string, content string, status Status) *Todo {
	return &Todo{
		ulid:    ulid,
		userId:  userUlid,
		title:   title,
		content: content,
		status:  status,
	}
}
