import request from "./request";

// test
export function testHttps() {
  return request({
    url: "/plugin_manage/info",
    method: "get",
  });
}