import { defineConfig } from 'vite'
import reactRefresh from '@vitejs/plugin-react-refresh'
import macrosPlugin from 'vite-plugin-babel-macros'
import legacy from '@vitejs/plugin-legacy'
// @ts-ignore
import { resolve } from 'path'


// https://vitejs.dev/config/
export default defineConfig({
  esbuild: {
    jsxFactory: 'jsx',
    jsxInject: 'import { jsx } from "@emotion/react"',
  },
  plugins: [reactRefresh({ exclude: [/node_modules/] }), macrosPlugin(), legacy({
    targets: ['chrome >= 70', 'firefox >= 72', 'edge >= 79', 'safari >= 13']
  })],
  resolve: {
    alias: {
      // @ts-ignore
      '@': resolve(__dirname, './src')
    }
  }
})
