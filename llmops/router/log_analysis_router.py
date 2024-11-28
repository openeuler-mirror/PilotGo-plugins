# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
from flask import Blueprint,request

from llmops.controller.log_analysis_controller import LogAnalysisController

log_analysis_router = Blueprint("log_analysis", __name__)
logAnalysisController = LogAnalysisController()


@log_analysis_router.route("/log_analysis", methods=["POST"])
def log_analysis():
    user_input = request.get_json("user_input")
    print(user_input)
    return logAnalysisController.log_analysis(user_input)
