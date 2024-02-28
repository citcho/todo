import { HeaderWithSidebar } from '@/components/templates/HeaderWithSidebar'
import { CompleteTodoList } from '@/components/organisms/Todo/CompleteTodoList'

export const CompleteList = () => {
  return (
    <HeaderWithSidebar>
      <CompleteTodoList />
    </HeaderWithSidebar>
  )
}
