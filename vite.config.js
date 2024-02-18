import { fileURLToPath, URL } from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({mode} ) => {
    // 加载所有变量
    const env = loadEnv(mode, process.cwd(), '');
    const PORT = env.PORT || 8071;

    return {
        plugins: [vue()],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        },
        server: {
            port: PORT,
            host: '127.0.0.1',
            proxy: {
                '/api': 'http://127.0.0.1:8072'
            },
        }
    }
})
