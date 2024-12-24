import requests

from llmops.config.config import init_config
from llmops.llm_tools.function_call.tools.os_tools import OsFunctionTools


class RpmFunctionTools:
    config = init_config()
    BASE_URL = config.app_conf.tool_baseurl

    @staticmethod
    def rpm_all(UUID: str):
        url = f"{OsFunctionTools.BASE_URL}/api/v1/rpm_all"
        params = {
            "uuid": UUID
        }
        response = requests.get(url, params=params)
        if response.status_code == 200:
            data = response.json()
            return data
        else:
            print(f"Failed to get rpm all. Status code: {response.status_code}")
            return None


    # @staticmethod
    # def rpm_source(UUID: str):
    #     url = f"{OsFunctionTools.BASE_URL}/api/v1/rpm_source"
    #     params = {
    #         "uuid": UUID
    #     }
    #     response = requests.get(url, params=params)
    #     if response.status_code == 200:
    #         data = response.json()
    #         return data
    #     else:
    #         print(f"Failed to get rpm source. Status code: {response.status_code}")
    #         return None

    @staticmethod
    def repos(UUID: str):
        url = f"{OsFunctionTools.BASE_URL}/api/v1/repos"
        params = {
            "uuid": UUID
        }
        response = requests.get(url, params=params)
        if response.status_code == 200:
            data = response.json()
            return data
        else:
            print(f"Failed to get repos. Status code: {response.status_code}")
            return None

    # @staticmethod
    # def rpm_info(UUID: str, package_name: str):
    #     url = f"{OsFunctionTools.BASE_URL}/api/v1/rpm_info"
    #     params = {
    #         "uuid": UUID,
    #         "package_name": package_name
    #     }
    #     response = requests.get(url, params=params)
    #     if response.status_code == 200:
    #         data = response.json()
    #         return data
    #     else:
    #         print(f"Failed to get rpm info. Status code: {response.status_code}")
    #         return None


rpm_tools_info = {
    "RpmAll": {
        "tool_name": "getRpmAll",
        "desc": (
            "此方法用于获取指定 UUID 客户端安装的所有 RPM 包列表。\n"
            "返回一个包含所有已安装 RPM 包的字符串列表，每个包的格式为：\n"
            "'包名-版本号-架构'，例如 'tar-1.35-2.oe2403.aarch64'。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose RPM package list is being queried."
            )
        },
        "func": RpmFunctionTools.rpm_all
    },


"   Repos": {
        "tool_name": "getRepos",
        "desc": (
            "此方法用于获取指定 UUID 客户端的仓库（repos）信息，包括多个仓库的配置。\n"
            "返回的数据包含多个仓库的信息，每个仓库包含以下字段：\n"
            "1. File：仓库文件名。\n"
            "2. ID：仓库 ID。\n"
            "3. Name：仓库名称。\n"
            "4. MirrorList：镜像列表（如果有）。\n"
            "5. BaseURL：仓库的基本 URL。\n"
            "6. MetaLink：仓库的元数据链接。\n"
            "7. MetadataExpire：元数据过期时间。\n"
            "8. GPGCheck：是否启用 GPG 检查。\n"
            "9. GPGKey：仓库的 GPG 密钥 URL。\n"
            "10. Enabled：仓库是否启用（1 表示启用，0 表示禁用）。\n"
            "如果方法调用失败，请告知用户错误原因。"
        ),
        "args": {
            "UUID": (
                "str",
                "[*Required] UUID string identifies the client whose repository information is being queried."
            )
        },
        "func": RpmFunctionTools.repos
    }
}

