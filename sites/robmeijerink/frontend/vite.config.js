import { defineConfig } from 'vite'
import path from 'path'

export default defineConfig({
    base: '/assets/dist/',

    build: {
        outDir: '../public/assets/dist',
        assetsDir: '',
        emptyOutDir: true,
        manifest: true,
        rollupOptions: {
            input: path.resolve(__dirname, 'src/scripts/app.js'),
            output: {
                entryFileNames: `[name]-[hash].js`,
                chunkFileNames: `[name]-[hash].js`,
                assetFileNames: `[name]-[hash].[ext]`
            },
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
})
