package initialize

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"monitor/internal/configs"
	"path"
	"time"
)

// 初始化日志方法
func initLogger() {
	if configs.GetAppConfig().GetDebug() {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.AddHook(getLsInitHook())
	log.SetReportCaller(true)
}

func getLsInitHook() *lfshook.LfsHook {
	lfHook := lfshook.NewHook(writerLog(), &log.TextFormatter{DisableColors: true})
	return lfHook
}

func writerLog() *rotatelogs.RotateLogs {
	logPath := configs.GetAppConfig().GetLogPath()
	var cstSh, _ = time.LoadLocation("Asia/ShangHai")
	fileSuffix := configs.GetAppConfig().GetName() + "_" + time.Now().In(cstSh).Format("20060102") + ".log"
	logFilePath := path.Join(logPath, fileSuffix)

	logger, err := rotatelogs.New(
		logFilePath,
		rotatelogs.WithRotationCount(8),           // 文件最大保存份数
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		panic(err)
	}
	return logger
}
