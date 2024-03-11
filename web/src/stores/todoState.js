import { atom, selector, useRecoilValue, useSetRecoilState } from 'recoil'

import { axios } from '@/libs/axiosConfig'
import { useNotification } from '@/hooks/useNotification'

const todosState = atom({
  key: 'todosState',
  default: [],
})

const completeTodoListState = selector({
  key: 'completeTodoListState',
  get: ({ get }) => {
    const todos = get(todosState)
    return todos.filter((todo) => todo.isComplete)
  },
})

const incompleteTodoListState = selector({
  key: 'incompleteTodoListState',
  get: ({ get }) => {
    const todos = get(todosState)
    return todos.filter((todo) => !todo.isComplete)
  },
})

const isErrorState = atom({
  key: 'isTodoErrorState',
  default: false,
})

export const useTodoState = () => {
  const incompleteTodos = useRecoilValue(incompleteTodoListState)
  const completeTodos = useRecoilValue(completeTodoListState)
  const isError = useRecoilValue(isErrorState)

  return {
    incompleteTodos,
    completeTodos,
    isError,
  }
}

export const useTodoMutators = () => {
  const { success, error } = useNotification()
  const setTodos = useSetRecoilState(todosState)
  const setIsError = useSetRecoilState(isErrorState)

  const getTodoList = () => {
    axios
      .get('/todos')
      .then(({ data }) => {
        data.todos
          ? setTodos(data.todos)
          : setTodos([])
      })
      .catch(() => {
        setIsError(true)
        error('エラーが発生しました。時間をおいて再度お試しください。')
      })
  }

  const createTodo = (data) => {
    return new Promise((resolve, reject) => {
      axios
        .post('/todos', data)
        .then(() => {
          success('Todoを作成しました。')
          resolve()
        })
        .catch((err) => {
          reject()
          if (!err || !err.data) {
            error('エラーが発生しました。時間をおいて再度お試しください。')
          }
        })
    })
  }

  const completeTodo = (id) => {
    return new Promise((resolve, reject) => {
      axios
        .patch(`/todos/${id}/complete`)
        .then(() => {
          success('Todoを完了しました')
          getTodoList()
          resolve()
        })
        .catch((err) => {
          reject()
          if (!err || !err.data) {
            error('エラーが発生しました。時間をおいて再度お試しください。')
          }
        })
    })
  }

  return {
    getTodoList,
    createTodo,
    completeTodo,
  }
}
