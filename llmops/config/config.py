import os
import yaml
from dataclasses import dataclass


@dataclass
class AppConf:
    server: str
    port: str
    debug: bool


class Config:
    filename = "../llm-ops.yaml"

    def __init__(self):
        self.app_conf: AppConf = None

    def load_config(self):
        current_dir = os.path.dirname(os.path.abspath(__file__))
        config_path = os.path.join(current_dir, self.filename)

        try:
            with open(config_path, "r") as f:
                conf = yaml.safe_load(f)
        except IOError as e:
            raise ValueError("Load llm-ops config file failed") from e

        self.app_conf = AppConf(**conf["app"])


def init_config() -> Config:
    conf = Config()
    conf.load_config()
    return conf
