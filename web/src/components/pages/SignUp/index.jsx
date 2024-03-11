import { useNavigate, NavLink } from 'react-router-dom'
import { useUserMutators } from '@/stores/userState'
import { useState, useCallback } from 'react'

import { signUpSchema } from '@/schemas/auth'

import style from './index.module.css'

export const SignUp = () => {
  const navigate = useNavigate()
  const { signUp } = useUserMutators()

  const [sign, setSign] = useState({
    name: '',
    email: '',
    password: '',
  })

  const [errors, setErrors] = useState({
    name: '',
    email: [],
    password: '',
  })

  const handleSubmit = (event) => {
    event.preventDefault()
    const { success, error } = signUpSchema.safeParse(sign)
    if (error) {
      setErrors(error.flatten().fieldErrors)
      return
    }
    if (success) {
      setErrors({ email: [], password: [] })
      signUp(sign.name, sign.email, sign.password).then(() => {
        navigate('/signin', { replace: true })
      })
    }
  }

  return (
    <div className='wrapper'>
      <div className={style['signup-card']}>
        <form onSubmit={handleSubmit} className={style['form']}>
          <div className={style['entry']}>
            <label htmlFor='name' className={style['label']}>
              Name
            </label>
            <input
              type='text'
              id='name'
              className={style['input']}
              onChange={useCallback(
                (event) => setSign({ ...sign, name: event.target.value }),
                [sign]
              )}
            />
            {errors.name && <p>{errors.name}</p>}
          </div>
          <div className={style['entry']}>
            <label htmlFor='email' className={style['label']}>
              Email
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
            {errors.password && <p>{errors.password}</p>}
          </div>
          <div className={style['submit']}>
            <button type='submit' className={style['submit-btn']}>
              Sign up
            </button>
          </div>
        </form>
        <div className={style['signin']}>
          <NavLink to='/signin' className={style['signin-link']}>
            Sign in
          </NavLink>
        </div>
      </div>
    </div>
  )
}
