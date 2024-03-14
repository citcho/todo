import path from 'path'
import { fileURLToPath } from 'url'

// eslint-disable-next-line import/no-extraneous-dependencies
import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

import basicSsl from '@vitejs/plugin-basic-ssl'

const fileName = fileURLToPath(import.meta.url)
const dirName = path.dirname(fileName)

export default defineConfig({
  envDir: `${dirName}/env`,
  plugins: [react(), basicSsl()],
  resolve: {
    alias: {
      '@/': `${dirName}/src/`,
    },
  },
  server: {
    host: 'dev-todo.citcho.com',
    open: true,
    port: 8000,
  },
})
