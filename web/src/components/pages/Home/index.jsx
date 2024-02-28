import { HeaderWithSidebar } from '@/components/templates/HeaderWithSidebar'
import { NavLink } from 'react-router-dom'

export const Home = () => {
  return (
    <HeaderWithSidebar>
      <p>HOME画面</p>
      <NavLink to='/login'>ログイン</NavLink>
    </HeaderWithSidebar>
  )
}
