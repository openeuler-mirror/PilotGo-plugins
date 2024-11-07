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
