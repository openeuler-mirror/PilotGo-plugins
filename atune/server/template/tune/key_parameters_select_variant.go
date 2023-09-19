package tune

type KeyParametersSelectVariantApp struct{}

func (kpsv *KeyParametersSelectVariantApp) Info() *TuneInfo {
	info := &TuneInfo{
		TuneName:      "key_parameters_select_variant",
		WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/key_parameters_select_variant.tar.gz /tmp/tune/ && cd /tmp/tune/ && tar -xzvf key_parameters_select_variant.tar.gz",
		Prepare:       "cd /tmp/tune/key_parameters_select_variant && sh prepare.sh",
		Tune:          "atune-adm tuning --project key_parameters_select_variant --detail key_parameters_select_variant_client.yaml",
		Restore:       "atune-adm tuning --restore --project key_parameters_select_variant",
	}
	return info
}
