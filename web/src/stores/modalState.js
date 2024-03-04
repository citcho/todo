import { useCallback } from 'react'
import { atom, useRecoilValue, useSetRecoilState } from 'recoil'

const modalState = atom({
  key: 'modalState',
  default: false,
})

/**
 * モーダルの開閉状態を管理するためのカスタムフック
 * @param {string} key ユニークなキー
 * @returns {object} useModalState, useModalMutators
 */
export const useModal = (key) => {
  /**
   * モーダルの開閉状態を返す
   * @returns {boolean} modalIsOpen モーダルの開閉状態
   */
  const useModalState = () => {
    const modalIsOpen = useRecoilValue(modalState)

    return { modalIsOpen }
  }

  /**
   * モーダルの開閉状態を変更する関数を返す
   * @returns {object} openModal, closeModal
   * @returns {Function} openModal モーダルを開く関数
   * @returns {Function} closeModal モーダルを閉じる関数
   */
  const useModalMutators = () => {
    const setModalState = useSetRecoilState(modalState)
    const openModal = useCallback(() => {
      return setModalState(true)
    }, [setModalState])

    const closeModal = useCallback(() => {
      return setModalState(false)
    }, [setModalState])

    return { openModal, closeModal }
  }

  return { useModalState, useModalMutators }
}
