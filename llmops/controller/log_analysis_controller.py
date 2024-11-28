# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
from flask import jsonify

from llmops.service.log_analysis_service import LogAnalysisService
from llmops.utils.response import success


class LogAnalysisController:
    def __init__(self):
        self.logAnalysis = LogAnalysisService()

    def log_analysis(self,user_input):
        log_analysis = self.logAnalysis.loganalysis(user_input)
        return success(log_analysis,"返回成功")