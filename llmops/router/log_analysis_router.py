# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
from Agently.Agent.Agent import Agent
from flask import Blueprint, request,current_app, Response

from llmops.controller.log_analysis_controller import LogAnalysisController
from llmops.llm_tools.log_analysis.log_analysis_workflow import generate
from llmops.utils.response import fail

aasf = "asdfsadf"
log_analysis_router = Blueprint("log_analysis", __name__)


@log_analysis_router.route("/log_analysis/stream", methods=["POST"])
def log_analysis_stream():
    # 从查询参数获取 user_input
    agent = current_app.config["AGENT"]
    user_input = request.form.get('user_input')
    print(user_input)
    if not user_input:
        return fail("error:user_input 参数是必需的")
    # 返回流式响应
    return Response(generate(agent, user_input),
                    content_type="text/event-stream",
                    status=200,
                    mimetype="text/event-stream",
                    headers={"Cache-Control": "no-cache", "Connection": "keep-alive"})


