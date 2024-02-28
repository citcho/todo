import { HeaderWithSidebar } from '@/components/templates/HeaderWithSidebar'
import { TodoList } from '@/components/organisms/Todo/TodoList'

export const List = () => {
  return (
    <HeaderWithSidebar>
      <TodoList />
    </HeaderWithSidebar>
  )
}
