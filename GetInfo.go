package getInfo

import (
	"github.com/zhangyiming748/getInfo/util"
	"github.com/zhangyiming748/log"
	"github.com/zhangyiming748/replace"
	"os/exec"
	"strconv"
)

/*
获取视频文件的帧数
*/
func GetVideoFrame(FullPath string) int {
	cmd := exec.Command("ffprobe", "-v", "error", "-count_frames", "-select_streams", "v:0", "-show_entries", "stream=nb_read_frames", "-of", "default=nokey=1:noprint_wrappers=1", FullPath)
	/*
		> -v error:这隐藏了"info"输出(版本信息等),使解析更容易.
		> -count_frames:计算每个流的帧数,并在相应的流部分中报告.
		> -select_streams v:0 :仅选择视频流.
		> -show_entries stream = nb_read_frames :只显示读取的帧数.
		> -of default = nokey = 1:noprint_wrappers = 1 :将输出格式(也称为"writer")设置为默认值,不打印每个字段的键(nokey = 1),不打印节头和页脚(noprint_wrappers = 1).
	*/
	log.Debug.Printf("生成的命令是:%s\n", cmd)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Warn.Panicf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Warn.Panicf("cmd.Run产生的错误:%v\n", err)
	}

	tmp := make([]byte, 1024)
	stdout.Read(tmp)
	t := string(tmp)
	t = replace.Replace(t)
	if frame, err := strconv.Atoi(t); err == nil {
		log.Debug.Printf("文件帧数约%d\n", frame)
		return frame
	}
	log.Warn.Println("读取文件帧数出错")
	return 0
}

/*
获取大于1080P的视频
*/
func GetOutOfFHD(dir, pattern string) {
	fs := util.GetMultiFiles(dir, pattern)
	for _, file := range fs {
		if util.BiggerThenFHD(file.FullPath) {
			log.Debug.Printf("大于1080P的文件:%v\n", file.FullPath)
		} else {
			log.Debug.Printf("小于1080P的文件:%v\n", file.FullPath)
		}
	}
}
