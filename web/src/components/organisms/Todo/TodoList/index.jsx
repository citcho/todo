import { memo, useCallback, useEffect, useMemo, useState } from 'react'
import { useTodoState, useTodoMutators } from '@/stores/todoState'
import { TodoTable } from '@/components/organisms/Todo/TodoTable'
import { useTodoList } from './index.hooks'
import { TodoModal } from '@/components/molecules/TodoModal'
import { useModal } from '@/stores/modalState'

export const TodoList = memo(() => {
  const { getTodoList, completeTodo } = useTodoMutators()
  const { todoRows } = useTodoList()
  const { incompleteTodos } = useTodoState()
  const memoizedTodoRows = useMemo(
    () => todoRows(incompleteTodos),
    [todoRows, incompleteTodos]
  )
  const [todo, setTodo] = useState({ id: '', content: '' })
  const { useModalState, useModalMutators } = useModal()
  const { modalIsOpen } = useModalState()
  const { openModal, closeModal } = useModalMutators()

  useEffect(() => {
    getTodoList()
  }, [])

  const handleRowClick = useCallback((id, content) => {
    setTodo({ id, content })
    openModal()
  })

  const handleCompleteIconClick = useCallback(
    (id) => {
      completeTodo(id)
    },
    [completeTodo]
  )

  return (
    <>
      {incompleteTodos.length ? (
        <>
          <TodoTable
            rows={memoizedTodoRows}
            onRowClick={handleRowClick}
            onCompleteIconClick={handleCompleteIconClick}
          />
          <TodoModal
            modalIsOpen={modalIsOpen}
            onRequestClose={closeModal}
            todo={todo}
          />
        </>
      ) : (
        <p>Todoはありません</p>
      )}
    </>
  )
})
