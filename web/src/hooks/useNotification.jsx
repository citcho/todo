import { toast } from 'react-toastify'

import styles from '@/components/molecules/Toast/index.module.css'

export const useNotification = () => {
  const success = (message, customId = undefined) => {
    toast.success(message, {
      position: toast.POSITION.TOP_CENTER,
      autoClose: 5000,
      closeOnClick: false,
      draggable: false,
      className: styles['toast-message-success'],
      toastId: customId,
    })
  }

  const error = (message, customId = undefined) => {
    toast.error(message, {
      autoClose: false,
      position: toast.POSITION.TOP_CENTER,
      hideProgressBar: true,
      closeOnClick: false,
      draggable: false,
      className: styles['toast-message-error'],
      toastId: customId,
    })
  }

  const info = (message, customId = undefined) => {
    toast.info(message, {
      autoClose: false,
      position: toast.POSITION.TOP_CENTER,
      hideProgressBar: true,
      closeOnClick: false,
      draggable: false,
      className: styles['toast-message-info'],
      toastId: customId,
    })
  }

  const warning = (message, customId = undefined) => {
    toast.warning(message, {
      autoClose: false,
      position: toast.POSITION.TOP_CENTER,
      hideProgressBar: true,
      closeOnClick: false,
      draggable: false,
      className: styles['toast-message-warning'],
      toastId: customId,
    })
  }

  return { success, error, info, warning }
}
