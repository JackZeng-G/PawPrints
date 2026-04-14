import axios from 'axios'
import { emitter } from '../utils/emitter'

const http = axios.create({
  baseURL: '/api',
  timeout: 10000,
})

http.interceptors.response.use(
  (res) => res,
  (err) => {
    const msg = err.response?.data?.error || err.message || '请求失败'
    emitter.emit('toast', { type: 'error', message: msg })
    return Promise.reject(err)
  }
)

export default http
