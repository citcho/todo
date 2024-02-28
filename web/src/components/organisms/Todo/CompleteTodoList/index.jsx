import { memo } from 'react'
import { TodoTable } from '@/components/organisms/Todo/TodoTable'

export const CompleteTodoList = memo(() => {
  return <TodoTable />
})
