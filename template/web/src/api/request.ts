import axios from 'axios';

// 1.创建axios实例
const request = axios.create({
  baseURL: '',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
  }
});

// 2.1添加请求拦截器
request.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  },
);

// 2.2添加响应拦截器
request.interceptors.response.use(
  (response: any) => {
    return response;
  },
  (error) => {
    if (error.response) {
      return Promise.reject(error.response.data);
    }
  },
);

export default request;
