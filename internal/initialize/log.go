package initialize

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"go-admin/internal/configs"
	utils "go-admin/internal/utils/tools"
	"os"
	"path"
	"time"
)

// 初始化日志方法
func initLogger() {

	if ok, _ := utils.PathExists(configs.GetLogConfig().GetLogPath()); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", configs.GetLogConfig().GetLogPath())
		// 创建日志文件夹
		_ = os.Mkdir(configs.GetLogConfig().GetLogPath(), os.ModePerm)
	}
	log.SetLevel(getConfigLogLevel())

	log.AddHook(getLsInitHook())

	log.SetReportCaller(configs.GetLogConfig().GetShowLine())
}

func getLsInitHook() *lfshook.LfsHook {
	lfHook := lfshook.NewHook(writerLog(), &log.TextFormatter{DisableColors: true})
	return lfHook
}

func writerLog() *rotatelogs.RotateLogs {
	logPath := configs.GetLogConfig().GetLogPath()
	var cstSh, _ = time.LoadLocation("Asia/ShangHai")
	fileSuffix := time.Now().In(cstSh).Format("20060102") + ".log"
	logFilePath := path.Join(logPath, fileSuffix)

	option := []rotatelogs.Option{
		rotatelogs.WithRotationCount(8),             // 文件最大保存份数
		rotatelogs.WithRotationTime(24 * time.Hour), // 日志切割时间间隔
	}
	if configs.GetLogConfig().GetLinkName() != "" {
		option = append(option, rotatelogs.WithLinkName(configs.GetLogConfig().GetLinkName()))
	}

	logger, err := rotatelogs.New(
		logFilePath,
		option...,
	)
	if err != nil {
		panic(err)
	}
	return logger
}

func getConfigLogLevel() log.Level {
	var level log.Level
	switch configs.GetLogConfig().GetLevel() { // 初始化配置文件的Level
	case "debug":
		level = log.DebugLevel
	case "info":
		level = log.InfoLevel
	case "warn":
		level = log.WarnLevel
	case "trace":
		level = log.TraceLevel
	case "error":
		level = log.ErrorLevel
	case "panic":
		level = log.PanicLevel
	case "fatal":
		level = log.FatalLevel
	default:
		level = log.InfoLevel
	}
	return level
}
