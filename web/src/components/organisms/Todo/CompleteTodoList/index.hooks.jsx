export const useTodoList = () => {
  const todoRows = (todos) =>
    todos.map((todo) => {
      const sanitizedTitle = todo.title.replace(
        /<("[^"]*"|'[^']*'|[^'">])*>/g,
        ''
      )
      const sanitizedContent = todo.content.replace(
        /<("[^"]*"|'[^']*'|[^'">])*>/g,
        ''
      )
      // 本文の文字を丸め込み（100文字以上の場合）
      const roundedContent =
        todo.content.length > 100
          ? `${sanitizedContent.substr(0, 100)}...`
          : sanitizedContent

      return {
        id: todo.id,
        title: sanitizedTitle,
        content: roundedContent,
        isComplete: todo.isComplete,
      }
    })

  return {
    todoRows,
  }
}
