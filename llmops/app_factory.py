from flask import Flask
from config.config import init_config


def create_app() -> Flask:
    app = Flask(__name__)

    # 初始化http服务
    config = init_config()
    app.config["SERVER"] = config.app_conf.server
    app.config["PORT"] = config.app_conf.port
    app.config["DEBUG"] = config.app_conf.debug

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
