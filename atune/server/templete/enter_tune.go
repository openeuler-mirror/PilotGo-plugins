package templete

import "openeuler.org/PilotGo/atune-plugin/templete/tune"

const (
	Compress                   = "compress"
	CompressExcept             = "compress_Except"
	Ffmpeg                     = "ffmpeg"
	Fio                        = "fio"
	GccCompile                 = "gcc_compile"
	GoGc                       = "go_gc"
	Graphicsmagick             = "graphicsmagick"
	Iozone                     = "iozone"
	KeyParametersSelect        = "key_parameters_select"
	KeyParametersSelectVariant = "key_parameters_select_variant"
	Mariadb                    = "mariadb"
	Memcached                  = "memcached"
	Memory                     = "memory"
	MysqlSysbench              = "mysql_sysbench"
	Nginx                      = "nginx"
	OpenGauss                  = "openGauss"
	Redis                      = "redis"
	Spark                      = "spark"
	TensorflowTrain            = "tensorflow_train"
	Tidb                       = "tidb"
	Tomcat                     = "tomcat"
)

func GetTuneInfo(tuneName string) *tune.TuneInfo {
	switch tuneName {
	case GccCompile:
		return tune.TuneGroupApp.GccCompile.Info()
	default:
		return nil
	}
}
func Prepare(tuneName string) error {
	switch tuneName {
	case GccCompile:
		return tune.TuneGroupApp.GccCompile.Prepare()
	default:
		return nil
	}
}
func StartTune(tuneName string) error {
	switch tuneName {
	case GccCompile:
		return tune.TuneGroupApp.GccCompile.StartTune()
	default:
		return nil
	}
}
func Restore(tuneName string) error {
	switch tuneName {
	case GccCompile:
		return tune.TuneGroupApp.GccCompile.Restore()
	default:
		return nil
	}
}
