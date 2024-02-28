import { createBrowserRouter } from 'react-router-dom'

import { Home } from '@/components/pages/Home'
import { Login } from '@/components/pages/Login'
import { Todo } from '@/components/pages/Todo'
import { List } from '@/components/pages/Todo/List'
import { CompleteList } from '@/components/pages/Todo/CompleteList'

export const router = (checkLogin) => {
  return createBrowserRouter([
    { path: '/', element: <Home />, loader: () => checkLogin() },
    {
      path: '/login',
      element: <Login />,
      loader: ({ request }) => {
        const currentPath = request.url.split(/(?=\/)/g).slice(-1)[0]
        return checkLogin('/', currentPath)
      },
    },
    {
      path: '/todos',
      element: <Todo />,
      loader: () => checkLogin(),
      children: [
        { index: true, element: <List /> },
        { path: 'complete', element: <CompleteList /> },
      ],
    },
  ])
}
