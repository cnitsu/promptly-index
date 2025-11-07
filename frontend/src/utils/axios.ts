import axios from "axios";

const instance = axios.create({
  timeout: 1000,
})

instance.defaults.baseURL = '/api'
instance.defaults.headers.common['Authorization'] = localStorage.getItem('user-token')

export default instance
