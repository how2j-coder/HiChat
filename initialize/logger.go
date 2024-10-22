package initialize

import (
	"HiChat/global"
	"HiChat/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

const blankStr = "  "

type Level = zapcore.Level

type LoggerWriter struct {
	test string
}

func (lw *LoggerWriter) Write(p []byte) (n int, err error) {
	splitStr := strings.Split(string(p), "--")
	logInfo := strings.Join(splitStr[:len(splitStr)-1], blankStr)
	logContent := splitStr[len(splitStr)-1]
	msg := fmt.Sprintf("%s\n%s", logInfo, logContent)
	return os.Stdout.Write([]byte(msg))
}

func (lw *LoggerWriter) Sync() error {
	return nil
}

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
	// 存储文件/控制
	var fileEncoder, consoleEncoder zapcore.Encoder

	// 调整编码器默认配置
	fileEncoderConfig := zap.NewProductionEncoderConfig()
	// 时间
	fileEncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05" + "]"))
	}
	// 日志等级
	fileEncoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("how2j" + "." + l.String())
	}
	// 调用行
	fileEncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(caller.TrimmedPath())
	}

	// 设置编码器
	if global.ServiceConfig.Log.Format == "json" {
		fileEncoder = zapcore.NewJSONEncoder(fileEncoderConfig)
	} else {
		fileEncoder = zapcore.NewConsoleEncoder(fileEncoderConfig)
	}
	// ----------------------------------------------------
	// 调整编码器默认配置
	logEncoderConfig := zapcore.EncoderConfig{
		TimeKey:          "Time",
		LevelKey:         "Level",
		NameKey:          "Logger",
		CallerKey:        "Caller",
		MessageKey:       "Message",
		StacktraceKey:    "StackTrace",
		LineEnding:       zapcore.DefaultLineEnding,
		FunctionKey:      zapcore.OmitKey,
		ConsoleSeparator: "--",
	}

	// 时间
	logEncoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("HI-CHAT")
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05" + "]"))
	}
	// 日志等级
	logEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 调用行
	logEncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(caller.FullPath())
	}

	// 设置编码器
	consoleEncoder = zapcore.NewConsoleEncoder(logEncoderConfig)
	writer := &LoggerWriter{}
	consoleCore := zapcore.NewCore(consoleEncoder, zapcore.AddSync(writer), level)

	return zapcore.NewTee(
		zapcore.NewCore(fileEncoder, getLogWriter(), level),
		consoleCore,
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
