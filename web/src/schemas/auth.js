import { z } from 'zod'

export const loginSchema = z.object({
  email: z
    .string()
    .email({ message: 'メールアドレスの形式が正しくありません' })
    .nonempty({ message: 'メールアドレスを入力してください' }),
  password: z.string().nonempty({ message: 'パスワードを入力してください' }),
})
