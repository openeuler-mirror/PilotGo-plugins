package service

import "openeuler.org/PilotGo/prometheus-plugin/global"

type PrometheusTarget struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID     string `json:"uuid"`
	TargetIP string `json:"targetIp"`
}

func GetPrometheusTarget() ([]string, error) {
	var ips []PrometheusTarget
	err := global.GlobalDB.Order("id desc").Find(&ips).Error
	if err != nil {
		return nil, err
	}
	var targets []string
	for _, ip := range ips {
		target := ip.TargetIP + ":9100"
		targets = append(targets, target)
	}
	return targets, nil
}
