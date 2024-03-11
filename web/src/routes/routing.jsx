import { createBrowserRouter } from 'react-router-dom'

import { Home } from '@/components/pages/Home'
import { SignUp } from '@/components/pages/SignUp'
import { SignIn } from '@/components/pages/SignIn'
import { Todo } from '@/components/pages/Todo'
import { List } from '@/components/pages/Todo/List'
import { CompleteList } from '@/components/pages/Todo/CompleteList'
import { Create } from '@/components/pages/Todo/Create'

export const router = (checkSignIn) => {
  return createBrowserRouter([
    { path: '/', element: <Home />, loader: () => checkSignIn() },
    {
      path: '/signup',
      element: <SignUp />,
    },
    {
      path: '/signin',
      element: <SignIn />,
      loader: ({ request }) => {
        const currentPath = request.url.split(/(?=\/)/g).slice(-1)[0]
        return checkSignIn('/', currentPath)
      },
    },
    {
      path: '/todos',
      element: <Todo />,
      loader: () => checkSignIn(),
      children: [
        { index: true, element: <List /> },
        { path: 'complete', element: <CompleteList /> },
        { path: 'create', element: <Create /> },
      ],
    },
  ])
}
