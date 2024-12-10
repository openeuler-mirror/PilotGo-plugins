# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
# 日志分析

from Agently.Agent.Agent import Agent

def generate(agent, user_input):
    # 假设 agent 的生成器返回逐步生成的数据
    generator = (
        agent
        .input(user_input)
        .info("用中文对系统日志进行解释，结果用字典形式例如：{reason:'解释'}")
        .instruct("目标语言：中文")
        .output({
            "analyse": ("str", "根据{user_input}出现的问题进行分析")
            # "reasons": [("str", "{user_input}根据找出报错的原因并按列表格式输出，添加编号")]
        })
        .get_instant_generator()
    )
    for chunk in generator:
        # 生成的数据以 'data: ' 开头，按格式发送
        # delta = chunk["delta"]
        # analyse = chunk["analyse"]
        yield f"data: {chunk}\n\n"

