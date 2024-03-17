import PropTypes from 'prop-types'
import { useCallback, useState } from 'react'

import { createTodoSchema } from '@/schemas/todo'
import { useTodoMutators } from '@/stores/todoState'

import styles from './index.module.css'

export const TodoCreate = ({ onSuccess }) => {
  const [todo, setTodo] = useState({ title: '', content: '' })
  const [errors, setErrors] = useState({ title: [], content: [] })
  const { createTodo } = useTodoMutators()

  const handleTitleInputChange = useCallback((event) => {
    setTodo((prev) => ({ ...prev, title: event.target.value }))
    if (event.target.value) setErrors((prev) => ({ ...prev, title: [] }))
  }, [])

  const handleContentInputChange = useCallback((event) => {
    setTodo((prev) => ({ ...prev, content: event.target.value }))
    if (event.target.value) setErrors((prev) => ({ ...prev, content: [] }))
  }, [])

  const handleSubmit = useCallback(
    (event) => {
      event.preventDefault()
      const { success, error } = createTodoSchema.safeParse(todo)
      if (error) {
        setErrors(error.flatten().fieldErrors)
        return
      }
      if (success) {
        setErrors({ name: [] })
        createTodo(todo)
          .then(() => onSuccess())
          .catch((err) => setErrors({ name: [err.message] }))
      }
    },
    [todo, createTodo, onSuccess]
  )

  return (
    <div className='wrapper'>
      <h2 className={styles['title']}>Create Todo</h2>
      <form className={styles['form']} onSubmit={handleSubmit}>
        <div className={styles['entry']}>
          <label htmlFor='content' className={styles['label']}>
            Title
          </label>
          <input
            type='content'
            id='content'
            className={styles['input']}
            onChange={handleTitleInputChange}
          />
          <p>{errors.title ? errors.title[0] : ''}</p>
        </div>
        <div className={styles['entry']}>
          <label htmlFor='content' className={styles['label']}>
            Content
          </label>
          <input
            type='content'
            id='content'
            className={styles['input']}
            onChange={handleContentInputChange}
          />
          <p>{errors.content ? errors.content[0] : ''}</p>
        </div>
        <div className={styles['submit']}>
          <button type='submit' className={styles['submit-btn']}>
            Submit
          </button>
        </div>
      </form>
    </div>
  )
}

TodoCreate.displayName = 'TodoCreate'
TodoCreate.propTypes = {
  onSuccess: PropTypes.func.isRequired,
}
