import Axios, { isAxiosError } from 'axios'
import Cookies from 'js-cookie'

const baseURL = import.meta.env.VITE_API_URL

export const axios = Axios.create({
  baseURL,
  withCredentials: true
})

axios.interceptors.request.use((request) => {
  const token = Cookies.get('token')
  if (token) {
    request.headers.Authorization = `Bearer ${token}`
  }
  return request
})

axios.interceptors.response.use(
  (response) => response,
  (error) => {
    if (isAxiosError(error)) {
      return Promise.reject(error.response)
    }
    return Promise.reject(error)
  }
)
