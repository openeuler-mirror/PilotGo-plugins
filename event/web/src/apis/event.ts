import request from "./index.ts";

export function getEventList(params?:any) {
  return request({
    url: "/plugin/event/query",
    method: "get",
    params
  })
}