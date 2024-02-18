import axios from 'axios'
import {ElMessage } from "element-plus";
import router from "@/router/index.js";
import { removeToken } from "@/utils/auth.js";

const service = axios.create({
    withCredentials: true,
    timeout: 10000
})

//请求配置
service.interceptors.request.use(
    config => {
        return config
    },
    error => {
        console.log(error)
        return Promise.reject(error)
    }
)

service.interceptors.response.use(
    response => {
        if (response.status !== 200) {
            console.log('error' + response)
            ElMessage ({
                message: "请求错误",
                type: 'error',
                duration: 5 * 1000
            })
        }
        return response

    },
    error => {
        console.log('err' + error)
        if (error.response.status === 401){
            removeToken()
            router.push('/login')
        }
        ElMessage.error(error.response.data)
        return Promise.reject(error)
    }
)

export default service