import axios from 'axios'


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
            Message({
                message: "请求错误",
                type: 'error',
                duration: 5 * 1000
            })
        }
    },
    error => {
        console.log('err' + error)
        Message({
            message: error.message,
            type: 'error',
            duration: 5 * 1000
        })
        return Promise.reject(error)
    }
)

export default service