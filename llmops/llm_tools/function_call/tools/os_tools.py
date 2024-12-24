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
            # 接口参数列表
            # code = data['code']
            # msg = data['msg']
            # ip = data['data']['ip']
            # department = data['data']['department']
            # state = data['data']['state']
            # platform = data['data']['platform']
            # platform_version = data['data']['platform_version']
            # kernel_arch = data['data']['kernel_arch']
            # kernel_version = data['data']['kernel_version']
            # cpu_num = data['data']['cpu_num']
            # model_name = data['data']['model_name']
            # memory_total = data['data']['memory_total']
            # disk_usage = data['data']['disk_usage']
            # immutable = data['data']['immutable']
            # for disk in disk_usage:
            #     print(f"  Device: {disk['device']}")
            #     print(f"  Path: {disk['path']}")
            #     print(f"  Total Space: {disk['total']}")
            #     print(f"  Used Percentage: {disk['used_percent']}%")

            return data
        else:
            print(f"Failed to get system info. Status code: {response.status_code}")
            return None
    # 获取代理列表
    @staticmethod
    def get_agent_list(UUID: str):
        # 定义目标URL
        print("OsFunctionTools be called!")
        url = f"{OsFunctionTools.BASE_URL}/api/v1/api/agent_list"

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
            # 遍历列表中的每个元素
            # for item in data:
            #     agent_uuid = item['agent_uuid']
            #     agent_version = item['agent_version']
            return data
        else:
            print(f"Failed to get agent_list info. Status code: {response.status_code}")
            return None
    # OS info
    @staticmethod
    def os_info(UUID: str):
        url = f"{OsFunctionTools.BASE_URL}/api/v1/api/os_info"
        params = {
            "uuid": UUID
        }
        response = requests.get(url, params=params)
        if response.status_code == 200:
            data = response.json()
            return data
        else:
            print(f"Failed to get os info. Status code: {response.status_code}")
            return None

    # CPU and Memory Information
    @staticmethod
    def cpu_info(UUID: str):
        url = f"{OsFunctionTools.BASE_URL}/api/v1/api/cpu_info"
        params = {
            "uuid": UUID
        }
        response = requests.get(url, params=params)
        if response.status_code == 200:
            data = response.json()
            return data
        else:
            print(f"Failed to get cpu info. Status code: {response.status_code}")
            return None


    @staticmethod
    def disk_use(UUID: str):
        url = f"{OsFunctionTools.BASE_URL}/api/v1/disk_use"
        params = {
            "uuid": UUID
        }
        response = requests.get(url, params=params)
        if response.status_code == 200:
            data = response.json()
            return data
        else:
            print(f"Failed to get disk use. Status code: {response.status_code}")
            return None

    # @staticmethod
    # def disk_info(UUID: str):
    #     url = f"{OsFunctionTools.BASE_URL}/api/v1/disk_info"
    #     params = {
    #         "uuid": UUID
    #     }
    #     response = requests.get(url, params=params)
    #     if response.status_code == 200:
    #         data = response.json()
    #         return data
    #     else:
    #         print(f"Failed to get disk info. Status code: {response.status_code}")
    #         return None


    # @staticmethod 可以用一个方法调用以上方法，但未知是否需要处理不同接口返回的数据
    # def fetch_data(endpoint: str, params=None):
    #     # 定义目标URL
    #     url = f"{OsFunctionTools.BASE_URL}{endpoint}"
    #
    #     # 发送GET请求
    #     response = requests.get(url, params=params)
    #
    #     # 检查响应状态码
    #     if response.status_code == 200:
    #         # 解析JSON响应
    #         data = response.json()
    #         return data
    #     else:
    #         print(f"Failed to fetch data. Status code: {response.status_code}")
    #         return None





# os function_call文字的描述。用来给llm识别各个方法的作用。
os_tools_info = {
    "GetFullSystemData": {
        "tool_name":"getAgentOverview",
        "desc": "The function Obtain all system information about a specified UUID agent node, including:"
                "1.ip地址，是当前节点agent的ip地址"
                "2.department:组织名，当前客户端节点所属的组织"
                "3.state:当前节点的状态，是否在线"
                "4.platform：平台类型"
                "5.platform_version:平台版本"
                "6.kernel_arch:内核"
                "7.kernel_version:内核版本"
                "8.cpu_num:cpu数量"
                "9.model_name:cpu的硬件信息"
                "10.memory_total:剩余内存memory空间"
                "11.dis_usage:硬盘使用情况"
                "如果调用方法失败，请将原因告知用户",
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string  is a string that identifies the agent node."
            )
        },
        "func": OsFunctionTools.get_agent_overview
    },
    "GetAgentList": {
        "tool_name":"getAgentOverview",
        "desc": "此方法是取指定UUID所有客户端节点（agent）的节点信息，包括："
                "1.agent_uuid：客户单节点的唯一字符串"
                "2.agent_version：客户端运行的版本"
                "如果调用方法失败，请将原因告知用户",
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string  is a string that identifies the client node."
            )
        },
        "func": OsFunctionTools.get_agent_list
    },
    "OsInfo": {
        "tool_name": "getOsInfo",
        "desc": (
            "此方法用于获取指定 UUID 客户端的操作系统信息，包括：\n"
            "1. IP：设备的 IP 地址。\n"
            "2. Platform：操作系统平台名称。\n"
            "3. PlatformVersion：操作系统平台版本。\n"
            "4. PrettyName：操作系统的友好显示名称。\n"
            "5. KernelVersion：内核版本号。\n"
            "6. KernelArch：内核的架构。\n"
            "7. HostId：主机唯一标识。\n"
            "8. Uptime：设备的启动时间。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose OS information is being queried."
            )
        },
        "func": OsFunctionTools.os_info
    },
    "CpuInfo": {
        "tool_name": "getCpuInfo",
        "desc": (
            "此方法用于获取指定 UUID 客户端的 CPU 信息，包括：\n"
            "1. ModelName：CPU 的型号名称。\n"
            "2. CpuNum：CPU 的核心数量。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose CPU information is being queried."
            )
        },
        "func": OsFunctionTools.cpu_info
    },
    "MemoryInfo": {
        "tool_name": "getMemoryInfo",
        "desc": (
            "此方法用于获取指定 UUID 客户端的内存信息，包括：\n"
            "1. MemTotal：总内存大小（单位：KB）。\n"
            "2. MemFree：空闲内存大小（单位：KB）。\n"
            "3. MemAvailable：可用内存大小（单位：KB）。\n"
            "4. Buffers：缓冲区大小（单位：KB）。\n"
            "5. Cached：缓存大小（单位：KB）。\n"
            "6. SwapCached：交换缓存大小（单位：KB）。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose memory information is being queried."
            )
        },
        "func": OsFunctionTools.memory_info
    },
    "DiskUse": {
        "tool_name": "getDiskUse",
        "desc": (
            "此方法用于获取指定 UUID 客户端的磁盘使用情况。\n"
            "返回的数据包含多个磁盘分区的使用情况，每个分区的信息包含以下字段：\n"
            "1. device：设备名称，例如 '/dev/dm-0'。\n"
            "2. path：挂载路径，例如 '/'。\n"
            "3. fstype：文件系统类型，例如 'ext2/ext3'。\n"
            "4. total：磁盘总容量，例如 '33G'。\n"
            "5. used：已使用的磁盘空间，例如 '8G'。\n"
            "6. usedPercent：已使用空间的百分比，例如 '26%'。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose disk usage information is being queried."
            )
        },
        "func": OsFunctionTools.disk_use
    }

}

