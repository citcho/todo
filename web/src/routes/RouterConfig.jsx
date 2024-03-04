import { RouterProvider } from 'react-router-dom'

import { useUserMutators } from '@/stores/userState'

import { router } from '@/routes/routing'

export const RouterConfig = () => {
  const { checkSignIn } = useUserMutators()
  return <RouterProvider router={router(checkSignIn)} />
}
