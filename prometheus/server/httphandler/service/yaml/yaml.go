package yaml

import (
	"errors"
	"os"

	"gitee.com/openeuler/PilotGo-plugins/sdk/utils/command"
	"gopkg.in/yaml.v2"
	"openeuler.org/PilotGo/prometheus-plugin/global"
)

func BackupPrometheusYML() error {
	cmd := "cp " + global.GlobalPrometheusYml + " " + global.GlobalPrometheusYml + ".bak"
	exitcode, _, stderr, err := command.RunCommand(cmd)
	if exitcode == 0 && stderr == "" && err == nil {
		return nil
	}
	return err
}

func UpdatePrometheusYML(rulePath string) error {
	YML := PrometheusYML{
		Global: struct {
			ScrapeInterval     string "yaml:\"scrape_interval\""
			EvaluationInterval string "yaml:\"evalution_interval\""
		}{
			ScrapeInterval:     "15s",
			EvaluationInterval: "15s"},
		RuleFiles: []string{rulePath},
		ScrapeConfigs: []struct {
			JobName       string "yaml:\"job_name\""
			HTTPSdConfigs []struct {
				Url             string "yaml:\"url\""
				RefreshInterval string "yaml:\"refresh_interval\""
			} "yaml:\"http_sd_configs\""
		}{
			{
				JobName: "node_exporter",
				HTTPSdConfigs: []struct {
					Url             string "yaml:\"url\""
					RefreshInterval string "yaml:\"refresh_interval\""
				}{
					{
						Url:             "http://192",
						RefreshInterval: "60s",
					},
				},
			},
		},
	}

	f, err := os.OpenFile(global.GlobalPrometheusYml, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	yaml.FutureLineWrap()
	encoder := yaml.NewEncoder(f)

	err = encoder.Encode(&YML)
	if err != nil {
		return err
	}
	return nil
}

func InitVerificationOrRollBack() (error, error) {
	cmd1 := "systemctl restart prometheus"
	exitcode1, _, stderr1, err1 := command.RunCommand(cmd1)
	if exitcode1 == 0 && stderr1 == "" && err1 == nil {
		return nil, nil
	}

	cmd2 := "cp " + global.GlobalPrometheusYml + ".bak" + " " + global.GlobalPrometheusYml
	exitcode2, _, stderr2, err2 := command.RunCommand(cmd2)
	if exitcode2 == 0 && stderr2 == "" && err2 == nil {
		return errors.New("There is an error in prometheus yml: %s" + err1.Error()), nil
	}

	return err1, errors.New("prometheus yml rollback has failed:%s" + err2.Error())
}
