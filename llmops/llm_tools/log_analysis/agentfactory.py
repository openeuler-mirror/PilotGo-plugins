# * Copyright (c) KylinSoft  Co., Ltd. 2024.All rights reserved.
# * PilotGo-plugins licensed under the Mulan Permissive Software License, Version 2. 
# * See LICENSE file for more details.
# * Author: yzy_dev <yuzhiyuan@kylinos.cn> 
# * Date: Thu Nov 28 15:00:41 2024 +0800 
import Agently

from llmops.config.config import Config


def agentfactory(config: Config):
    agent_factory = (
        Agently.AgentFactory()
        .set_settings("current_model", config.llm_conf.current_model)
        # .set_settings("model.OAIClient.auth.api_key", config.llm_conf.apikey)
        .set_settings("model.OAIClient.options.model", config.llm_conf.model)
        .set_settings("model.OAIClient.url", config.llm_conf.baseurl)
    )
    return agent_factory

