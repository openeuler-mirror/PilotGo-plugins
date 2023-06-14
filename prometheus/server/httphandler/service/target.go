package service

import (
	"errors"

	"openeuler.org/PilotGo/prometheus-plugin/global"
)

type PrometheusTarget struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UUID     string `json:"uuid"`
	TargetIP string `json:"targetIp"`
	Port     string `json:"port"`
	ID_idx   string `gorm:"index"`
}

func GetPrometheusTarget() ([]string, error) {
	var ips []PrometheusTarget
	err := global.GlobalDB.Raw("SELECT * FROM prometheus_target ORDER BY id DESC").Scan(&ips).Error
	if err != nil {
		return []string{}, err
	}

	if len(ips) == 0 {
		return []string{}, errors.New("ip targets is null")
	}
	var targets []string
	for _, ip := range ips {
		target := ip.TargetIP + ":" + ip.Port
		targets = append(targets, target)
	}
	return targets, nil
}

func AddPrometheusTarget(pt PrometheusTarget) error {
	t := PrometheusTarget{
		UUID:     pt.UUID,
		TargetIP: pt.TargetIP,
		Port:     pt.Port,
	}
	err := global.GlobalDB.Save(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePrometheusTarget(pt PrometheusTarget) error {
	var t PrometheusTarget
	err := global.GlobalDB.Where("uuid = ?", pt.UUID).Unscoped().Delete(t).Error
	if err != nil {
		return err
	}
	return nil
}
