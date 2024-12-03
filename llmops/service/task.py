# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: zhanghan2021 <zhanghan@kylinos.cn> 
# * Date: Thu Nov 7 16:30:24 2024 +0800 
import uuid
import time
from threading import Thread

# 内存中任务存储结构
tasks = {}


# 定义任务类
class Task:
    def __init__(self, user_id, input):
        self.task_id = str(uuid.uuid4())
        self.user_id = user_id
        self.current_step = 1
        self.step_results = {}
        self.input = input
        self.status = "等待中"  # 等待中, 正在处理, 等待确认, 完成
        self.options = []  # LLM 生成的方案

    def update_step_result(self, step, result):
        self.step_results[step] = result

    def update_status(self, status):
        self.status = status

    def set_options(self, options):
        self.options = options


# 创建新任务
def create_new_task(user_id, input):
    task = Task(user_id, input)
    tasks[task.task_id] = task

    # 启动一个线程来异步处理 LLM 任务
    thread = Thread(target=llm_process, args=(task.task_id,))
    thread.start()

    return task.task_id


# 获取任务状态
def get_task_status(task_id):
    task = tasks.get(task_id)
    if not task:
        return None

    return {
        "task_id": task.task_id,
        "status": task.status,
        "current_step": task.current_step,
        "step_results": task.step_results,
        "options": task.options,
    }


# 用户确认方案，继续任务
def confirm_task_option(task_id, selected_option):
    task = tasks.get(task_id)
    if not task:
        return {"error": "未找到该任务"}

    if task.status != "等待确认":
        return {"error": "该任务无等待人工确认的操作"}

    # 用户选择方案后继续任务
    task.update_step_result(task.current_step, f"Confirmed {selected_option}")
    task.current_step += 1  # 进入下一步
    task.update_status("正在处理")

    # 启动一个线程继续执行任务
    thread = Thread(target=continue_task, args=(task.task_id, selected_option))
    thread.start()

    return {
        "message": "Option confirmed, task continued",
        "task_id": task_id,
        "status": "正在处理",
    }


# 模拟 LLM 处理
def llm_process(task_id):
    task = tasks.get(task_id)
    if not task:
        return

    task.update_status("正在处理")
    time.sleep(2)  # 模拟 LLM 生成方案的延迟

    # 假设 LLM 生成了 3 个方案
    options = [f"方案 {i}" for i in range(1, 4)]
    task.set_options(options)
    task.update_step_result(task.current_step, options)
    task.update_status("等待确认")


# 模拟继续执行任务
def continue_task(task_id, selected_option):
    task = tasks.get(task_id)
    if not task:
        return

    task.update_status("正在处理")
    time.sleep(2)  # 模拟方案执行的延迟

    # 模拟方案执行后的结果
    final_result = f"Final result for {selected_option}"
    task.current_step += 1
    task.update_step_result(task.current_step, final_result)
    task.update_status("完成")
