# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 15:01:23 2024 +0800 
from flask import Flask
from config.config import init_config
from llmops.router.log_analysis_router import log_analysis_router
from llmops.utils.agentfactoryUtils import agentfactory
from utils.logger import setup_logger
from router.task import task_blueprint


def create_app() -> Flask:
    app = Flask(__name__)

    # 初始化http服务
    config = init_config()
    app.config["SERVER"] = config.app_conf.server
    app.config["PORT"] = config.app_conf.port
    app.config["DEBUG"] = config.app_conf.debug

    # 设置日志
    logger = setup_logger()
    app.logger = logger

    # 注册蓝图
    app.register_blueprint(task_blueprint, url_prefix="/task")
    app.register_blueprint(log_analysis_router, url_prefix="/log")

    # agent对象存储在app.config内
    factory = agentfactory(config)  # 初始化agentfactory工厂类
    agent = factory.create_agent();  # 初始化agent
    app.config["AGENT"] = agent

    return app


def run_app():
    app = create_app()
    try:
        app.run(
            host=app.config["SERVER"],
            port=int(app.config["PORT"]),
            debug=app.config["DEBUG"],
        )
    except ValueError as e:
        app.logger.error(f"Invalid app config: {e}")
        exit(1)
