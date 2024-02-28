import { memo } from 'react'

import styles from './index.module.css'
import { NavLink } from 'react-router-dom'

export const Sidebar = memo(() => {
  return (
    <aside className={styles['sidebar']}>
      <ul className={styles['list']}>
        <li className={styles['item']}>
          <NavLink to='/' className={styles['link']}>
            HOME
          </NavLink>
        </li>
        <li className={styles['item']}>
          <NavLink to='/todos' className={styles['link']}>
            Yet
          </NavLink>
        </li>
        <li className={styles['item']}>
          <NavLink to='/todos/complete' className={styles['link']}>
            Done
          </NavLink>
        </li>
      </ul>
    </aside>
  )
})

Sidebar.displayName = 'TheSidebar'
