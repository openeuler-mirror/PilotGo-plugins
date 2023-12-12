import request from './request';

// 获取结果分页
export function getResultLists(data: object) {
    return request({
      url: '/plugin/atune/results',
      method: 'get',
      params: data,
    });
  }

// 高级搜索结果分页
export function searchResult(data: object) {
    return request({
      url: '/plugin/atune/result_search',
      method: 'get',
      params: data,
    });
  }

  // 删除结果
export function deleteResult(data: object) {
  return request({
    url: '/plugin/atune/result_delete',
    method: 'delete',
    data,
  });
}