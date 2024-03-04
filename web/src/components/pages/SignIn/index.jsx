import { useNavigate } from 'react-router-dom'
import { useUserMutators } from '@/stores/userState'
import { useState, useCallback } from 'react'

import { signSchema } from '@/schemas/auth'

import style from './index.module.css'

export const SignIn = () => {
  const navigate = useNavigate()
  const { signIn } = useUserMutators()

  const [sign, setSign] = useState({
    email: '',
    password: '',
  })

  const [errors, setErrors] = useState({
    email: [],
    password: [],
  })

  const handleSubmit = (event) => {
    event.preventDefault()
    const { success, error } = signSchema.safeParse(sign)
    if (error) {
      setErrors(error.flatten().fieldErrors)
      return
    }
    if (success) {
      setErrors({ email: [], password: [] })
      signIn(sign.email, sign.password).then(() => {
        navigate('/', { replace: true })
      })
    }
  }

  return (
    <div className='wrapper'>
      <div className={style['signin-card']}>
        <h2 className={style['title']}>Sign in to your account</h2>
        <form onSubmit={handleSubmit} className={style['form']}>
          <div className={style['entry']}>
            <label htmlFor='email' className={style['label']}>
              Your email
            </label>
            <input
              type='email'
              id='email'
              className={style['input']}
              onChange={useCallback(
                (event) => setSign({ ...sign, email: event.target.value }),
                [sign]
              )}
            />
            {errors.email[0] && <p>{errors.email[0]}</p>}
          </div>
          <div className={style['entry']}>
            <label htmlFor='password' className={style['label']}>
              Password
            </label>
            <input
              type='password'
              id='password'
              className={style['input']}
              onChange={useCallback(
                (event) => setSign({ ...sign, password: event.target.value }),
                [sign]
              )}
            />
            {errors.password[0] && <p>{errors.password[0]}</p>}
          </div>
          <div className={style['submit']}>
            <button type='submit' className={style['submit-btn']}>
              Sign in
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}
