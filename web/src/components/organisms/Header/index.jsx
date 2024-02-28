import { memo } from 'react'

import styles from './index.module.css'
import { NavLink } from 'react-router-dom'

export const Header = memo(() => {
  return (
    <header className={styles['header']}>
      <div className={styles['inner']}>
        <NavLink to='/' className={styles['logo']}>
          Todo App
        </NavLink>
      </div>
    </header>
  )
})
Header.displayName = 'Header'
