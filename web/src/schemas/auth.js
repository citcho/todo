import { z } from 'zod'

export const signSchema = z.object({
  email: z
    .string()
    .email({ message: 'メールアドレスの形式が正しくありません' })
    .min(1, { message: 'メールアドレスを入力してください' }),
  password: z.string().min(1, { message: 'パスワードを入力してください' }),
})
