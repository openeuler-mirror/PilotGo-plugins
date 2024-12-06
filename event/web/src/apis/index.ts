/* 
 * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
 * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
 * See LICENSE file for more details.
 * Author: 赵振芳 <zhaozhenfang@kylinos.cn>
 * Date: Mon Nov 25 17:29:28 2024 +0800
 */
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
  baseURL: '', 
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
    return response; 
  },
  (error) => {
    return error
  }
);

export default function request(config: AxiosRequestConfig) {
  return instance(config);
}