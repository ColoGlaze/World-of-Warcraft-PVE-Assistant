import axios from 'axios'

// 创建axios实例
const request = axios.create({
    baseURL: '/api',
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json'
    }
})

// 请求拦截器
request.interceptors.request.use(
    (config) => {
        // 在发送请求之前做些什么
        // console.log('请求发送:', config.url)

        // 可以在这里添加token等认证信息
        // const token = localStorage.getItem('token')
        // if (token) {
        //   config.headers.Authorization = `Bearer ${token}`
        // }

        return config
    },
    (error) => {
        // console.error('请求错误:', error)
        return Promise.reject(error)
    }
)

// 响应拦截器
request.interceptors.response.use(
    (response) => {
        // console.log('响应接收:', response.data)
        return response.data
    },
    (error) => {
        // console.error('响应错误:', error)

        // 统一错误处理
        let message = '网络错误'
        if (error.response) {
            switch (error.response.status) {
                case 400:
                    message = '请求参数错误'
                    break
                case 401:
                    message = '未授权访问'
                    break
                case 403:
                    message = '禁止访问'
                    break
                case 404:
                    message = '请求资源不存在'
                    break
                case 500:
                    message = '服务器内部错误'
                    break
                default:
                    message = error.response.data?.message || '请求失败'
            }
        } else if (error.request) {
            message = '网络连接失败'
        }

        return Promise.reject(new Error(message))
    }
)

export default request