import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig(({ command }) => ({
    base: command === 'serve' ? '/' : '/assets/rob/',

    build: {
        outDir: '../public/dist',
        emptyOutDir: true,
        manifest: true,
        rollupOptions: {
            input: path.resolve(__dirname, 'src/scripts/app.js'),
        },
    },

    server: {
        host: '127.0.0.1',
        port: 5173,
        strictPort: true,
        origin: 'http://localhost:5173',
        cors: true,
        allowedHosts: ['robmeijerink.test'],

        hmr: {
            host: 'localhost',
            protocol: 'ws',
        },

        watch: {
            ignored: ['!**/views/**'],
        },
    },
}))
