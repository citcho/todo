import clsx from 'clsx'
import PropTypes from 'prop-types'
import { memo } from 'react'

import { Header } from '@/components/organisms/Header'
import { Sidebar } from '@/components/organisms/Sidebar'

import styles from './index.module.css'

export const HeaderWithSidebar = memo(({ children, isAllowOverflowScroll }) => {
  return (
    <div className={styles.layout}>
      <Header />
      <div className={styles.wrapper}>
        <Sidebar />
        <main
          className={clsx(styles.main, {
            [styles['overflow-scroll']]: isAllowOverflowScroll,
          })}
        >
          {children}
        </main>
      </div>
    </div>
  )
})

HeaderWithSidebar.displayName = 'HeaderWithSidebar'
HeaderWithSidebar.propTypes = {
  children: PropTypes.oneOfType([PropTypes.node]).isRequired,
  isAllowOverflowScroll: PropTypes.bool,
}
HeaderWithSidebar.defaultProps = {
  isAllowOverflowScroll: false,
}
