import request from "./request";

// 获取所有的任务列表
export function getTaskLists(data: Object) {
  return request({
    url: "/plugin/atune/tasks",
    method: "get",
    params: data,
  });
}

// 高级搜索任务列表分页
export function searchTask(data: object) {
  return request({
    url: "/plugin/atune/task_search",
    method: "get",
    params: data,
  });
}

// 新增任务
export function addTask(data: object) {
  return request({
    url: "/plugin/atune/run",
    method: "post",
    data,
  });
}

// 删除任务
export function deleteTask(data: object) {
  return request({
    url: "/plugin/atune/task_delete",
    method: "delete",
    data,
  });
}

// 获取所有的调优列表
export function getAtuneAllName() {
  return request({
    url: "/plugin/atune/all",
    method: "get",
  });
}

// 获取某个调优对象的具体信息
export function getAtuneInfo(data: object) {
  return request({
    url: "/plugin/atune/info",
    method: "get",
    params: data,
  });
}

// 获取调优模板列表分页
export function getTuneLists(data: object) {
  return request({
    url: "/plugin/atune/tunes",
    method: "get",
    params: data,
  });
}

// 保存调优模板
export function saveTune(data: object) {
  return request({
    url: "/plugin/atune/save_tune",
    method: "post",
    data,
  });
}

// 删除调优模板
export function deleteTune(data: object) {
  return request({
    url: "/plugin/atune/delete_tune",
    method: "delete",
    data,
  });
}

// 高级搜索模板列表分页
export function searchTune(data: object) {
  return request({
    url: "/plugin/atune/search_tune",
    method: "get",
    params: data,
  });
}

// 编辑模板更新
export function updateTune(data: object) {
  return request({
    url: "/plugin/atune/update_tune",
    method: "post",
    data,
  });
}
