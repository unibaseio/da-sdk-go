package log

import (
	"fmt"
	"sync"

	"github.com/unibaseio/da-sdk-go/lib/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var mLogger *zap.SugaredLogger
var mLoglevel zap.AtomicLevel
var lk sync.Mutex

func Logger(name string) *zap.SugaredLogger {
	lk.Lock()
	defer lk.Unlock()
	return mLogger.Named(name)
}

// StartLogger starts
func init() {
	mLoglevel = zap.NewAtomicLevel()

	// Logs go to stderr (Unix convention: diagnostics on stderr, data on stdout).
	// This keeps CLI/tool stdout clean for machine parsing; node logs under
	// systemd/journald capture stderr identically. Override with LOG_FILE.
	outputs := []string{"stderr"}
	debugWriter, _, err := zap.Open(outputs...)
	if err != nil {
		panic(fmt.Sprintf("unable to open logging output: %v", err))
	}

	lf := env.Str(env.LogFile, "")
	if lf != "" {
		debugWriter = getLogWriter(lf)
	}

	encoder := getEncoder()

	core := zapcore.NewCore(encoder, debugWriter, mLoglevel)

	// NewProduction
	logger := zap.New(core, zap.AddCaller())

	mLogger = logger.Sugar()

	l := getLogLevel(env.Str(env.LogLevel, ""))

	mLoglevel.SetLevel(l)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "sub",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename + ".log",
		MaxSize:    100, //MB
		MaxBackups: 3,
		MaxAge:     30, //days
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLogLevel(level string) zapcore.Level {
	l := zapcore.InfoLevel
	switch level {
	case "debug", "DEBUG":
		l = zapcore.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		l = zapcore.InfoLevel
	case "warn", "WARN":
		l = zapcore.WarnLevel
	case "error", "ERROR":
		l = zapcore.ErrorLevel
	case "dpanic", "DPANIC":
		l = zapcore.DPanicLevel
	case "panic", "PANIC":
		l = zapcore.PanicLevel
	case "fatal", "FATAL":
		l = zapcore.FatalLevel

	}

	return l
}

func SetLogLevel(level string) error {
	l := getLogLevel(level)

	lk.Lock()
	defer lk.Unlock()

	mLoglevel.SetLevel(l)
	return nil
}
