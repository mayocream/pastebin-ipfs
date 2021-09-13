import { defineConfig } from 'vite'
import reactRefresh from '@vitejs/plugin-react-refresh'
import macrosPlugin from 'vite-plugin-babel-macros'
// @ts-ignore
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  esbuild: {
    jsxFactory: 'jsx',
    jsxInject: 'import { jsx } from "@emotion/react"',
  },
  plugins: [reactRefresh(), macrosPlugin()],
  resolve: {
    alias: {
      // @ts-ignore
      '@': resolve(__dirname, './src')
    }
  }
})
