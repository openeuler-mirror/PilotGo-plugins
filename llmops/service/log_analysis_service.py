# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800
from Agently.Agent.Agent import Agent

from llmops.llm_tools.log_analysis import log_analysis_workflow


class LogAnalysisService:
    def loganalysis(self, userinput: str, agent: Agent):
        result = log_analysis_workflow.logworkflow(userinput, agent)
        return result
