package logger

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const LogPath = "./output/logs/"
const FileSuffix = ".log"

var Log = logrus.New()

// 文件日志分割；使用：logger.Log.Error("msg")
func New() *logrus.Logger {
	setLogger(Log, LogPath)
	Log.SetLevel(logrus.InfoLevel)
	return Log
}

// 设置日志中间件
func setLogger(log *logrus.Logger, logPath string) {
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(logPath, "debug"),
		logrus.InfoLevel:  writer(logPath, "info"),
		logrus.WarnLevel:  writer(logPath, "warn"),
		logrus.ErrorLevel: writer(logPath, "error"),
		logrus.FatalLevel: writer(logPath, "fatal"),
		logrus.PanicLevel: writer(logPath, "panic"),
	}, &Formatter{})

	log.Hooks.Add(lfHook)
}

// 文件设置
func writer(logPath string, level string) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, level)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fileSuffix := time.Now().In(cstSh).Format("2006-01-02") + FileSuffix

	logier, err := rotatelogs.New(
		logFullPath+"-"+fileSuffix,
		rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Hour*24*15),    // 最长保留时间15天
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}
