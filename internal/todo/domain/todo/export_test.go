package todo

func NewTodoFixture(id string, userId string, title string, content string, isComplete bool) *Todo {
	return &Todo{
		id:         id,
		userId:     userId,
		title:      title,
		content:    content,
		isComplete: isComplete,
	}
}
