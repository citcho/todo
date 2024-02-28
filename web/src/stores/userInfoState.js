import { redirect } from 'react-router-dom'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil'

import { axios } from '@/libs/axiosConfig'

import { useNotification } from '@/hooks/useNotification'

const userInfoState = atom({
  key: 'userInfoState',
  default: {
    name: '',
    email: '',
  },
})

const isLoadingState = atom({
  key: 'isLoadingState',
  default: false,
})

/**
 * ユーザー情報とローディングフラグを返す
 * @returns {object} userInfo, isLoading
 * @returns {object} userInfo ユーザー情報
 * @returns {boolean} isLoading ローディングフラグ
 */
export const useUserInfoState = () => {
  const userInfo = useRecoilValue(userInfoState)
  const isLoading = useRecoilValue(isLoadingState)

  return { userInfo, isLoading }
}

export const useUserInfoMutators = () => {
  const { error } = useNotification()
  const setUserInfo = useSetRecoilState(userInfoState)
  const setIsLoading = useSetRecoilState(isLoadingState)

  /**
   * ログインチェック
   * @param {string} redirectPath リダイレクト先のパス
   * @param {string} currentPath 現在のパス
   * @returns {Promise<Function | null>} リダイレクト関数 or null
   */
  const checkLogin = (redirectPath, currentPath) => {
    return new Promise((resolve, reject) => {
      axios
        .get('/me')
        .then(({ data }) => {
          if (data.user) {
            setUserInfo(data.user)
            if (redirectPath) {
              resolve(redirect(redirectPath))
            } else {
              resolve(null)
            }
          } else if (currentPath === '/login') {
            resolve(null)
          } else {
            resolve(redirect('/login'))
          }
        })
        .catch((err) => {
          switch (err.status) {
            case 401:
              if (currentPath === '/login') {
                resolve(null)
              } else {
                resolve(redirect('/login'))
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

  /**
   * ログイン処理
   * @param {string} email メールアドレス
   * @param {string} password パスワード
   * @returns {Promise<void>} Promise
   */
  const login = (email, password) => {
    return new Promise((resolve) => {
      setIsLoading(true)
      axios
        .post('/login', {
          email,
          password,
        })
        .then(() => {
          checkLogin()
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
        .finally(() => {
          setIsLoading(false)
        })
    })
  }

  return {
    checkLogin,
    login,
  }
}
