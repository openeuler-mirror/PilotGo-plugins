# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 16:30:24 2024 +0800 
from flask import Blueprint
from controller.task import create_task, task_status, confirm_option

task_blueprint = Blueprint("tasks", __name__)

# 路由
task_blueprint.route("/create", methods=["POST"])(create_task)
task_blueprint.route("/status/<task_id>", methods=["GET"])(task_status)
task_blueprint.route("/confirm", methods=["POST"])(confirm_option)
