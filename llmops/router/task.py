from flask import Blueprint
from controller.task import create_task, task_status, confirm_option

task_blueprint = Blueprint("tasks", __name__)

# 路由
task_blueprint.route("/create", methods=["POST"])(create_task)
task_blueprint.route("/status/<task_id>", methods=["GET"])(task_status)
task_blueprint.route("/confirm", methods=["POST"])(confirm_option)
