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

def createagent(agent_factory: Agently.AgentFactory):
     agent = agent_factory.create_agent();
     return agent

