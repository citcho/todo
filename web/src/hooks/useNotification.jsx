import { toast } from 'react-toastify'

import styles from '@/components/molecules/Toast/index.module.css'

/**
 * 通知を表示する関数を返すhooks
 * @returns {object} success, error, info, warning
 */
export const useNotification = () => {
  /**
   * 成功の通知を表示する
   * @param {string} message 表示する成功メッセージ
   * @param {string | number} customId 各トーストを識別するためのID
   */
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

  /**
   * エラーの通知を表示する
   * @param {string} message 表示するエラーメッセージ
   * @param {string | number} customId 各トーストを識別するためのID
   */
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

  /**
   * 情報の通知を表示する
   * @param {string} message 表示する情報メッセージ
   * @param {string | number} customId 各トーストを識別するためのID
   */
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

  /**
   * 警告の通知を表示する
   * @param {string} message 表示する警告メッセージ
   * @param {string | number} customId 各トーストを識別するためのID
   */
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
