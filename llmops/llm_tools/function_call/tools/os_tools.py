import requests

from llmops.config.config import init_config


class OsFunctionTools:
    config = init_config()
    BASE_URL = config.app_conf.tool_baseurl

    @staticmethod
    def get_agent_overview(UUID:str):
        # 定义目标URL
        print("OsFunctionTools be called!")
        url = f"{OsFunctionTools.BASE_URL}/api/v1/api/agent_overview"

        # 如果需要传递查询参数，可以使用params参数
        params = {
            "uuid": UUID
            }

        # 发送GET请求
        response = requests.get(url, params=params)

        # 检查响应状态码
        if response.status_code == 200:
            # 解析JSON响应
            data = response.json()
            # print("Memory Info:", data)
            return data
        else:
            print(f"Failed to get memory info. Status code: {response.status_code}")
            return None


tools = {
    "GetFullSystemData": {
        "tool_name":"getAgentOverview",
        "desc": "获取指定UUID客户端节点的全部系统信息,包括："
                "1.ip地址，是当前节点的ip地址"
                "2.department:组织名，当前客户端节点所属的组织"
                "3.state:当前节点的状态，是否在线"
                "4.platform：平台类型"
                "5.platform_version:平台版本"
                "6.kernel_arch:内核"
                "7.kernel_version:内核版本"
                "8.cpu_num:cpu数量"
                "9.model_name:cpu的硬件信息"
                "10.memory_total:剩余内存memory空间"
                "11.dis_usage:硬盘使用情况",
        "args": {
            "UUID": (
                "str",
                "[*Required] Timezone string used in pytz.timezone() in Python"
            )
        },
        "func": OsFunctionTools.get_agent_overview
    },

}

