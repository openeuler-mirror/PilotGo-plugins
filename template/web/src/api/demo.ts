import request from "./request";

export function demoAPI(data?:Object) {
  return request({
    url: "/plugin/template/api/hello_world",
    method: "get",
    params:data
  })
}