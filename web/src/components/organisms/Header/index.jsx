import { memo } from 'react'
import { useNavigate, NavLink } from 'react-router-dom'

import { useUserMutators } from '@/stores/userState'

import styles from './index.module.css'

export const Header = memo(() => {
  const navigate = useNavigate()
  const { signOut } = useUserMutators()
  const handleSignOutButtonClick = () => {
    signOut().then(() => {
      navigate('/signin', { replace: true })
    })
  }

  return (
    <header className={styles['header']}>
      <div className={styles['inner']}>
        <NavLink to='/' className={styles['logo']}>
          Todo App
        </NavLink>
        <button
          type='submit'
          className={styles['signout-btn']}
          onClick={handleSignOutButtonClick}
        >
          Sign out
        </button>
        <></>
      </div>
    </header>
  )
})

Header.displayName = 'Header'
