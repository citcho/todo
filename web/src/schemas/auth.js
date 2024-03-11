import { z } from 'zod'

const signUpSchema = z.object({
  name: z
    .string()
    .min(1, { message: '名前を入力してください' }),
  email: z
    .string()
    .email({ message: 'メールアドレスの形式が正しくありません' })
    .min(1, { message: 'メールアドレスを入力してください' }),
  password: z.string().min(1, { message: 'パスワードを入力してください' }),
})

const signInSchema = z.object({
  email: z
    .string()
    .email({ message: 'メールアドレスの形式が正しくありません' })
    .min(1, { message: 'メールアドレスを入力してください' }),
  password: z.string().min(1, { message: 'パスワードを入力してください' }),
})

export { signUpSchema, signInSchema }