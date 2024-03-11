import { redirect } from 'react-router-dom'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil'

import { axios } from '@/libs/axiosConfig'
import { useNotification } from '@/hooks/useNotification'

const userState = atom({
  key: 'userState',
  default: {
    name: '',
    email: '',
  },
})

export const useUserState = () => {
  const user = useRecoilValue(userState)

  return { user }
}

export const useUserMutators = () => {
  const { success, error } = useNotification()
  const setUser = useSetRecoilState(userState)

  const checkSignIn = (redirectPath, currentPath) => {
    return new Promise((resolve, reject) => {
      axios
        .get('/me')
        .then(({ data }) => {
          if (data.user) {
            setUser(data.user)
            if (redirectPath) {
              resolve(redirect(redirectPath))
            } else {
              resolve(null)
            }
          } else if (currentPath === '/signin') {
            resolve(null)
          } else {
            resolve(redirect('/signin'))
          }
        })
        .catch((err) => {
          switch (err.status) {
            case 401:
              if (currentPath === '/signin') {
                resolve(null)
              } else {
                resolve(redirect('/signin'))
              }
              break
            default:
              reject(
                new Error(
                  'エラーが発生しました。時間をおいて再度お試しください。'
                )
              )
              break
          }
          reject(
            new Error('エラーが発生しました。時間をおいて再度お試しください。')
          )
        })
    })
  }

  const signIn = (email, password) => {
    return new Promise((resolve) => {
      axios
        .post('/signin', {
          email,
          password,
        })
        .then(() => {
          checkSignIn()
            .then(() => {
              resolve()
            })
            .catch((err) => {
              error(err.message)
            })
        })
        .catch(() => {
          error('メールアドレスまたはパスワードが間違っています。')
        })
    })
  }

  const signUp = (name, email, password) => {
    return new Promise((resolve) => {
      axios
        .post('/signup', {
          name,
          email,
          password,
        })
        .then(() => {
          success('ユーザー登録が完了しました。')
          resolve()
        })
        .catch(({ data }) => {
          error(data.message)
        })
    })
  }

  return {
    checkSignIn,
    signUp,
    signIn,
  }
}
