package tune

var TuneGroupApp = new(TuneGroup)

type TuneGroup struct {
	Compress                   CompressApp
	CompressExcept             CompressExceptApp
	Ffmpeg                     FfmpegApp
	Fio                        FioApp
	GccCompile                 GccComplieApp
	GoGc                       GoGcApp
	Graphicsmagick             GraphicsmagickApp
	Iozone                     IozoneApp
	KeyParametersSelect        KeyParametersSelectApp
	KeyParametersSelectVariant KeyParametersSelectVariantApp
	Mariadb                    MariadbApp
	Memcached                  MemcachedApp
	Memory                     MemoryApp
	MysqlSysbench              MysqlSysbenchApp
	Nginx                      NginxApp
	OpenGauss                  OpenGaussApp
	Redis                      RedisApp
	Spark                      SparkApp
	TensorflowTrain            TensorflowTrainApp
	Tidb                       TidbApp
	Tomcat                     TomcatApp
}

type TuneInfo struct {
	Prepare string `json:"prepare"`
	Tune    string `json:"tune"`
	Restore string `json:"restore"`
}