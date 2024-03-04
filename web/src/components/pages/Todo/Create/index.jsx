import { useCallback } from 'react'
import { useNavigate } from 'react-router-dom'

import { HeaderWithSidebar } from '@/components/templates/HeaderWithSidebar'
import { TodoCreate } from '@/components/organisms/Todo/TodoCreate'

export const Create = () => {
  const navigate = useNavigate()

  const handleCreateTodoSuccess = useCallback(() => {
    navigate('/todos')
  }, [navigate])
  return (
    <HeaderWithSidebar>
      <TodoCreate onSuccess={handleCreateTodoSuccess} />
    </HeaderWithSidebar>
  )
}
