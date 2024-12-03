# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 16:30:24 2024 +0800 
from flask import request
from utils.response import success, fail, not_found
from service.task import create_new_task, get_task_status, confirm_task_option


# 创建新任务
def create_task():
    data = request.json
    user_id = data.get("user_id")
    input = data.get("input")

    if not user_id or not input:
        return not_found(None, "用户id和输入内容不能为空")

    task_id = create_new_task(user_id, input)
    return success({"task_id": task_id, "status": "等待中"}, "创建异步任务成功")


# 查询任务状态
def task_status(task_id):
    task_info = get_task_status(task_id)
    if not task_info:
        return not_found(None, "未找到该任务")

    return success(task_info, "获取到任务状态")


# 确认用户选择的方案
def confirm_option():
    data = request.json
    task_id = data.get("task_id")
    selected_option = data.get("selected_option")

    if not task_id or not selected_option:
        return not_found(None, "任务id和所选方案内容不能为空")

    result = confirm_task_option(task_id, selected_option)
    if "error" in result:
        return fail(None, result)
    return success(result, "已确认所选方案")
