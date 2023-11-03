import request from './request';

// 获取所有的调优列表
export function getAtuneAllName() {
  return request({
    url: '/plugin/atune/all',
    method: 'get',
  });
}

// 获取某个调优对象的具体信息
export function getAtuneInfo(data: object) {
  return request({
    url: '/plugin/atune/info',
    method: 'get',
    params: data,
  });
}
