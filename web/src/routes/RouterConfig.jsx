import { RouterProvider } from 'react-router-dom'

import { useUserInfoMutators } from '@/stores/userInfoState'

import { router } from '@/routes/routing'

export const RouterConfig = () => {
  const { checkLogin } = useUserInfoMutators()
  return <RouterProvider router={router(checkLogin)} />
}
