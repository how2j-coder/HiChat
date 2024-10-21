package initialize

import (
	"HiChat/global"
	"HiChat/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Level = zapcore.Level

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

func InitLogger() {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()

	if global.ServiceConfig.Log.ShowLine {
		options = append(options, zap.AddCaller())
	}
	global.Logger = zap.New(getZapCore(), options...)
}

func createRootDir() {
	if ok, _ := utils.PathExists(global.ServiceConfig.Log.RootDir); !ok {
		_ = os.Mkdir(global.ServiceConfig.Log.RootDir, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.ServiceConfig.Log.Level {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

// 扩展 Zap
func getZapCore() zapcore.Core {
	var fileEncoder, consoleEncoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05" + "]"))
	}
	// 日志等级
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("how2j" + "." + l.String())
	}

	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(caller.TrimmedPath())
	}

	// 设置编码器
	if global.ServiceConfig.Log.Format == "json" {
		fileEncoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		fileEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	consoleEncoder = zapcore.NewConsoleEncoder(encoderConfig)
	//var writes = []zapcore.WriteSyncer{getLogWriter(), zapcore.AddSync(os.Stdout)}
	return zapcore.NewTee(
		zapcore.NewCore(fileEncoder, getLogWriter(), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)
}

// 使用 lumberjack 作为日志写入器
func getLogWriter() zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		Filename:   global.ServiceConfig.Log.RootDir + "/" + global.ServiceConfig.Log.Filename,
		MaxSize:    global.ServiceConfig.Log.MaxSize,
		MaxBackups: global.ServiceConfig.Log.MaxBackups,
		MaxAge:     global.ServiceConfig.Log.MaxAge,
		Compress:   global.ServiceConfig.Log.Compress,
	}

	return zapcore.AddSync(file)
}
