package logger

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

var (
	defaultLevel    = "debug" // output log levels debug, info, warn, error, default is debug
	defaultEncoding = formatConsole
	defaultIsSave   = false // False:output to terminal, true:output to file, default is false.

	defaultFilename      = "out.log" // File name
	defaultMaxSize       = 10        // Maximum file size (MB)
	defaultMaxBackups    = 100       // Maximum amount old files
	defaultMaxAge        = 30        // Maximum amount days for old documents
	defaultIsCompression = false     // Whether to compress and archive old files
	defaultIsLocalTime   = true      // Whether to use local time
)

type options struct {
	level    string
	encoding string
	isSave   bool
	fileConfig *fileOptions
	hooks []func(zapcore.Entry) error
}

type fileOptions struct {
	filename      string
	maxSize       int
	maxBackups    int
	maxAge        int
	isCompression bool
	isLocalTime   bool
}

func defaultOptions() *options {
	return &options{
		level:    defaultLevel,
		encoding: defaultEncoding,
		isSave:   false,
	}
}

// Option set the logger options.
type Option func(*options)
func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithLevel setting the log level
func WithLevel(levelName string) Option {
	return func(o *options) {
		levelName = strings.ToUpper(levelName)
		switch levelName {
		case levelDebug, levelInfo, levelWarn, levelError:
			o.level = levelName
		default:
			o.level = levelDebug
		}
	}
}

// WithFormat set the output log format, console or json
func WithFormat(format string) Option {
	return func(o *options) {
		if strings.ToLower(format) == formatJSON {
			o.encoding = formatJSON
		}
	}
}