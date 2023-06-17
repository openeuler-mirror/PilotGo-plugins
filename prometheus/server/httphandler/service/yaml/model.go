package yaml

type PrometheusYML struct {
	Global struct {
		ScrapeInterval     string `yaml:"scrape_interval"`
		EvaluationInterval string `yaml:"evalution_interval"`
	} `yaml:"global"`
	RuleFiles     []string `yaml:"rule_files"`
	ScrapeConfigs []struct {
		JobName       string `yaml:"job_name"`
		HTTPSdConfigs []struct {
			Url             string `yaml:"url"`
			RefreshInterval string `yaml:"refresh_interval"`
		} `yaml:"http_sd_configs"`
	} `yaml:"scrape_configs"`
}
