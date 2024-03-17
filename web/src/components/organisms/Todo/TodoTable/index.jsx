import PropTypes from 'prop-types'
import { memo } from 'react'
import { GoCheckCircleFill } from 'react-icons/go'
import styles from './index.module.css'

import tableStyles from '@/styles/table.module.css'

export const TodoTable = memo(({ rows, onRowClick, onCompleteIconClick }) => {
  const columns = [
    { id: 1, header: 'タイトル', size: 'sm' },
    { id: 2, header: '内容', size: 'sm' },
    { id: 3, header: '', size: 'auto' },
  ]

  return (
    <table className={tableStyles.table}>
      <thead>
        <tr>
          {columns.map((column) => (
            <th key={column.id} className={styles[`size--${column.size}`]}>
              {column.header}
            </th>
          ))}
        </tr>
      </thead>
      <tbody>
        {rows.map((row) => (
          <tr key={row.id} onClick={() => onRowClick(row.id, row.content)}>
            <td>{row.title}</td>
            <td>{row.roundedContent}</td>
            <td className={styles['align-left']}>
              {!row.isComplete && (
                <button
                  className={styles['complete-button']}
                  onClick={(event) => {
                    event.stopPropagation()
                    onCompleteIconClick(row.id)
                  }}
                >
                  <GoCheckCircleFill />
                </button>
              )}
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  )
})

TodoTable.displayName = 'TodoTable'
TodoTable.propTypes = {
  rows: PropTypes.arrayOf(
    PropTypes.exact({
      id: PropTypes.string.isRequired,
      title: PropTypes.string.isRequired,
      content: PropTypes.string.isRequired,
      roundedContent: PropTypes.string.isRequired,
      isComplete: PropTypes.bool.isRequired,
    })
  ).isRequired,
  onRowClick: PropTypes.func.isRequired,
  onCompleteIconClick: PropTypes.func,
}
