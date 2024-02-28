import { useNavigate } from 'react-router-dom'
import { useUserInfoMutators } from '@/stores/userInfoState'
import { useState, useCallback } from 'react'

import { loginSchema } from '@/schemas/auth'

import style from './index.module.css'

export const Login = () => {
  const navigate = useNavigate()
  const { login } = useUserInfoMutators()

  const [loginInfo, setLoginInfo] = useState({
    email: '',
    password: '',
  })

  const [errors, setErrors] = useState({
    email: [],
    password: [],
  })

  const handleSubmit = (event) => {
    event.preventDefault()
    const { success, error } = loginSchema.safeParse(loginInfo)
    if (error) {
      setErrors(error.flatten().fieldErrors)
      return
    }
    if (success) {
      setErrors({ email: [], password: [] })
      login(loginInfo.email, loginInfo.password).then(() => {
        navigate('/', { replace: true })
      })
    }
  }

  return (
    <div className='wrapper'>
      <div className={style['login-card']}>
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
                (event) =>
                  setLoginInfo({ ...loginInfo, email: event.target.value }),
                [loginInfo]
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
                (event) =>
                  setLoginInfo({ ...loginInfo, password: event.target.value }),
                [loginInfo]
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
