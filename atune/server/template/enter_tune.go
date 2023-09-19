package template

import "openeuler.org/PilotGo/atune-plugin/template/tune"

const (
	Compress                   = "compress"
	CompressExcept             = "compress_Except"
	Ffmpeg                     = "ffmpeg"
	Fio                        = "fio"
	GccCompile                 = "gcc_compile"
	GoGc                       = "go_gc"
	Graphicsmagick             = "graphicsmagick"
	Iozone                     = "iozone"
	Kafka                      = "kafka"
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

func GetTuneInfo(tuneName string) interface{} {
	switch tuneName {
	case Compress:
		return tune.TuneGroupApp.Compress.Info()
	case CompressExcept:
		return tune.TuneGroupApp.CompressExcept.Info()
	case Ffmpeg:
		return tune.TuneGroupApp.Ffmpeg.Info()
	case Fio:
		return tune.TuneGroupApp.Fio.Info()
	case GccCompile:
		return tune.TuneGroupApp.GccCompile.Info()
	case GoGc:
		return tune.TuneGroupApp.GoGc.Info()
	case Graphicsmagick:
		return tune.TuneGroupApp.Graphicsmagick.Info()
	case Iozone:
		return tune.TuneGroupApp.Iozone.Info()
	case Kafka:
		return tune.TuneGroupApp.Kafka.Info()
	case KeyParametersSelect:
		return tune.TuneGroupApp.KeyParametersSelect.Info()
	case KeyParametersSelectVariant:
		return tune.TuneGroupApp.KeyParametersSelectVariant.Info()
	case Mariadb:
		return tune.TuneGroupApp.Mariadb.Info()
	case Memcached:
		return tune.TuneGroupApp.Memcached.Info()
	case Memory:
		return tune.TuneGroupApp.Memory.Info()
	case MysqlSysbench:
		return tune.TuneGroupApp.MysqlSysbench.Info()
	case Nginx:
		return tune.TuneGroupApp.Nginx.Info()
	case OpenGauss:
		return tune.TuneGroupApp.OpenGauss.Info()
	case Redis:
		return tune.TuneGroupApp.Redis.Info()
	case Spark:
		return tune.TuneGroupApp.Spark.Info()
	case TensorflowTrain:
		return tune.TuneGroupApp.TensorflowTrain.Info()
	case Tidb:
		return tune.TuneGroupApp.Tidb.Info()
	case Tomcat:
		return tune.TuneGroupApp.Tomcat.Info()
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
