import { memo, useCallback, useEffect, useMemo, useState } from 'react'
import { useTodoState, useTodoMutators } from '@/stores/todoState'
import { TodoTable } from '@/components/organisms/Todo/TodoTable'
import { useTodoList } from './index.hooks'
import { TodoModal } from '@/components/molecules/TodoModal'
import { useModal } from '@/stores/modalState'

export const CompleteTodoList = memo(() => {
  const { getTodoList } = useTodoMutators()
  const { todoRows } = useTodoList()
  const { completeTodos } = useTodoState()
  const memoizedTodoRows = useMemo(
    () => todoRows(completeTodos),
    [todoRows, completeTodos]
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

  return (
    <>
      {completeTodos.length ? (
        <>
          <TodoTable rows={memoizedTodoRows} onRowClick={handleRowClick} />
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
