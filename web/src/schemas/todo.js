import { z } from 'zod'

export const createTodoSchema = z.object({
  title: z
    .string()
    .min(1, { message: 'タイトルを入力してください' })
    .max(255, { message: 'タイトルは255文字以内で入力してください' }),

  content: z
    .string()
    .max(1000, { message: '内容は1000文字以内で入力してください' }),
})
