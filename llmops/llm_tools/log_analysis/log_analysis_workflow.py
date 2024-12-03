# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
# 日志分析
import os

import Agently
from Agently.Agent.Agent import Agent

from llmops.config.config import init_config
from llmops.llm_tools.log_analysis.agentfactory import agentfactory


def logworkflow(user_input,agent: Agent):
    main_workflow = Agently.Workflow()
    log_agent = agent

    # @main_workflow.chunk("user_input")
    # def user_input(inputs, storage):
    #     storage.set(
    #         "user_input",
    #         # input("[请输入你的问题]: ")
    #         user_input
    #     )

    # 返回shell脚本
    @main_workflow.chunk("assistant_reply")
    def assistant_shell_reply(inputs, storage):
        print("分析中.....")
        # assistant_reason_reply = (
        #     agent
        #     .info("回答尽量精简")
        #     .user_info("用户是服务器运维人员，操作系统是macOs")
        #     .input({
        #         "user_input": storage.get("user_input"),
        #     })
        #     .instruct([
        #         "replay中输出语言为中文",
        #         f"请根据：{user_input}找出日志的问题",
        #         "用简洁的方式说明问题"
        #         "操作系统是maoOs"
        #     ])
        #     .output({
        #         "replay": ("str", "对{user_input}进行解释"),
        #         "improvement": ([("str", "根据用户输入{input}输出相应脚本处理用户的要求"), "找出改进方案"])
        #     })
        #     .start_async()
        # )
        # storage.set("replay", assistant_reason_reply.get("replay"))
        # print("*" * 50)
        # print("一.大模型分析：")
        # print(assistant_reason_reply.get("replay"))
        # print("*" * 50)
        # print("二.生成的脚本内容:")
        # storage.set("shell", assistant_reason_reply.get("improvement"))
        # print(assistant_reason_reply.get("improvement"))
        # print("*" * 50)
        # return assistant_reason_reply.get("improvement")
        assistant_reason_reply = (
            log_agent
            .input(user_input)
            .info("用中文对系统日志进行解释，结果用字典形式例如：{reason:'解释'}")
            .instruct("目标语言：中文")
            .output({
                "info_list": [
                    {
                        "知识对象": ("str", "分析{input}问题时，需要了解相关知识的具体对象"),
                        "关键知识点": ("str", "分析{input}问题时，需要了解的关键知识")
                    }
                ],
                "analyse": ("str", "根据{info_list}出现的问题进行解释和分析"),
                "reasons": [("str", "{info_list}根据找出报错的原因并按列表格式输出，添加编号")]
            })
            .start()
        )
        return assistant_reason_reply

    (
        main_workflow
        # .connect_to("user_input")
        .connect_to("assistant_reply")
        .connect_to("END")
    )
    # print(main_workflow.draw())
    result = main_workflow.start()
    return result


# r = logworkflow("Can not connect to MySQL server. Too many connections -mysql 1040错误")
# print(r['default'])
