# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 15:01:23 2024 +0800 
import os
import yaml
from dataclasses import dataclass


@dataclass
class AppConf:
    server: str
    port: str
    debug: bool

@dataclass
class LlmConf:
    model: str
    # apikey: str
    baseurl: str
    current_model: str


class Config:
    filename = "llm-ops.yaml"

    def __init__(self):
        self.app_conf: AppConf = None
        self.llm_conf: LlmConf = None
    def load_config(self):
        config_path = os.path.join(os.path.dirname(__file__), Config.filename)

        try:
            with open(config_path, "r") as f:
                conf = yaml.safe_load(f)
        except IOError as e:
            raise ValueError("Load llm-ops config file failed") from e

        self.app_conf = AppConf(**conf["app"])
        self.llm_conf = LlmConf(**conf["llm"])

def init_config() -> Config:
    """initialize llm-ops config"""
    conf = Config()
    conf.load_config()
    return conf
