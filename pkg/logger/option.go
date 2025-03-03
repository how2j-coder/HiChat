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

// ----------------- console -----------------
type options struct {
	level    string
	encoding string
	isSave   bool
	fileConfig *fileOptions
	hooks []func(zapcore.Entry) error
}

func defaultOptions() *options {
	return &options{
		level:    defaultLevel,
		encoding: defaultEncoding,
		isSave:   defaultIsSave,
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

// ----------------- file -----------------
type fileOptions struct {
	filename      string
	maxSize       int
	maxBackups    int
	maxAge        int
	isCompression bool
	isLocalTime   bool
}

func defaultFileOptions() *fileOptions {
	return &fileOptions{
		filename:      defaultFilename,
		maxSize:       defaultMaxSize,
		maxBackups:    defaultMaxBackups,
		maxAge:        defaultMaxAge,
		isCompression: defaultIsCompression,
		isLocalTime:   defaultIsLocalTime,
	}
}

// FileOption set the file options.
type FileOption func(*fileOptions)

func (o *fileOptions) apply(opts ...FileOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithFileName set log filename
func WithFileName(filename string) FileOption {
	return func(f *fileOptions) {
		if filename != "" {
			f.filename = filename
		}
	}
}

// WithFileMaxSize set maximum file size (MB)
func WithFileMaxSize(maxSize int) FileOption {
	return func(f *fileOptions) {
		if maxSize > 0 {
			f.maxSize = maxSize
		}
	}
}

// WithFileMaxBackups set maximum amount old files
func WithFileMaxBackups(maxBackups int) FileOption {
	return func(f *fileOptions) {
		if f.maxBackups > 0 {
			f.maxBackups = maxBackups
		}
	}
}

// WithFileMaxAge set maximum amount days for old documents
func WithFileMaxAge(maxAge int) FileOption {
	return func(f *fileOptions) {
		if f.maxAge > 0 {
			f.maxAge = maxAge
		}
	}
}

// WithFileIsCompression set whether to compress log files
func WithFileIsCompression(isCompression bool) FileOption {
	return func(f *fileOptions) {
		f.isCompression = isCompression
	}
}

// WithLocalTime set whether to use local time
func WithLocalTime(isLocalTime bool) FileOption {
	return func(f *fileOptions) {
		f.isLocalTime = isLocalTime
	}
}