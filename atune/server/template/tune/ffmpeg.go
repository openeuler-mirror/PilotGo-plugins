package tune

type FfmpegApp struct{}

type FfmpegImp struct {
	BaseTune TuneInfo
	Notes    string `json:"note"`
}

func (f *FfmpegApp) Info() *FfmpegImp {
	info := &FfmpegImp{
		BaseTune: TuneInfo{
			TuneName:      "ffmpeg",
			WorkDirectory: "mkdir -p /tmp/tune/ && cp -r ../../templete/ffmpeg.tar.gz /tmp/tune/ && cd /tmp/tune && tar -xzvf ffmpeg.tar.gz",
			Prepare:       "cd /tmp/tune/ffmpeg && sh prepare.sh",
			Tune:          "atune-adm tuning --project ffmpeg --detail ffmpeg_client.yaml",
			Restore:       "atune-adm tuning --restore --project ffmpeg",
		},
		Notes: "请先找一个测试视频将其命名为test.flv(基本参数,41m34s,640x352,flv,83.6MB,视频编码AVC1,音频编码AAC)，并将其放在/tmp/tune/ffmpeg目录下",
	}
	return info
}
