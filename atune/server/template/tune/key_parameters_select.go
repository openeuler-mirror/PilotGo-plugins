package tune

type KeyParametersSelectApp struct{}

func (kps *KeyParametersSelectApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "key_parameters_select",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/key_parameters_select.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf key_parameters_select.tar.gz",
		Prepare:       "cd /tmp/tune/key_parameters_select && sh prepare.sh",
		Tune:          "atune-adm tuning --project key_parameters_select --detail key_parameters_select_client.yaml",
		Restore:       "atune-adm tuning --restore --project key_parameters_select",
	}
	return info
}
