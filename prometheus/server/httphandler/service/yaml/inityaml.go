package yaml

import (
	"openeuler.org/PilotGo/prometheus-plugin/config"
)

func InitPrometheusYML(conf *config.Prometheus) error {
	if err := BackupPrometheusYML(); err != nil {
		return err
	}

	if err := UpdatePrometheusYML(conf.AlertYaml); err != nil {
		return err
	}

	updateErr, rollebackErr := InitVerificationOrRollBack()
	if updateErr != nil {
		return updateErr
	}
	if rollebackErr != nil {
		return rollebackErr
	}

	return nil
}
