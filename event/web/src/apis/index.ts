import axios, { type AxiosRequestConfig } from 'axios';

// 公共定义
export const RespCodeOK = 200
export interface RespInterface {
  code?: number;
  data?: any[];
  msg?: string;
  ok?:boolean;
  page?:number;
  size?: number;
  total?: number;
}
  

// 创建实例
const instance = axios.create({
  baseURL: '/', 
  timeout: 10000, 
});

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    return response.data; 
  },
  (error) => {
    return error
  }
);

export default function request(config: AxiosRequestConfig) {
  return instance(config);
}