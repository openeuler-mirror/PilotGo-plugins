package tune

type KafkaApp struct{}

type KafkaImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (k *KafkaApp) Info() *KafkaImp {
	info := &KafkaImp{
		BaseTune: TuneInfo{
			TuneName:      "kafka",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/kafka.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf kafka.tar.gz",
			Prepare:       "cd /tmp/tune/kafka && sh prepare.sh",
			Tune:          "atune-adm tuning --project kafka --detail ./kafka_client.yaml",
			Restore:       "atune-adm tuning --restore --project kafka",
		},
		Notes: "使用kafka atune之前, 请先配置好kafka环境",
	}
	return info
}
