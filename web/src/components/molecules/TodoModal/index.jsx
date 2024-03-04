import PropTypes from 'prop-types'
import { memo } from 'react'

import Modal from 'react-modal'

import styles from './index.module.css'

export const TodoModal = memo(({ modalIsOpen, onRequestClose, todo }) => {
  Modal.setAppElement(document.getElementById('root'))
  return (
    <Modal
      isOpen={modalIsOpen}
      onRequestClose={onRequestClose}
      closeTimeoutMS={300}
      shouldCloseOnEsc
      shouldCloseOnOverlayClick
      shouldFocusAfterRender
      shouldReturnFocusAfterClose
      className={styles.modal}
      overlayClassName={{
        base: styles['modal-mask'],
        afterOpen: styles['modal-mask--after-open'],
        beforeClose: styles['modal-mask--before-close'],
      }}
    >
      <div className={styles['modal-heading']}>{todo.content}</div>
    </Modal>
  )
})

TodoModal.displayName = 'TodoModal'
TodoModal.propTypes = {
  modalIsOpen: PropTypes.bool.isRequired,
  onRequestClose: PropTypes.func.isRequired,
  todo: PropTypes.exact({
    id: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
  }).isRequired,
}
