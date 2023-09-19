package tune

type TensorflowTrainApp struct{}

func (tensor *TensorflowTrainApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "tensorflow_train",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/tensorflow_train.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf tensorflow_train.tar.gz",
		Prepare:       "cd /tmp/tune/tensorflow_train && sh prepare.sh",
		Tune:          "atune-adm tuning --project tensorflow_train --detail tensorflow_train_client.yaml",
		Restore:       "atune-adm tuning --restore --project tensorflow_train",
	}
	return info
}
