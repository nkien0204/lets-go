package log

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

var one sync.Once

//func initZapLogger() *zap.Logger {
//	one.Do(func() {
//		var err error
//		var cfg zap.Config
//		mode := viper.GetString("mode")
//		if mode == "prod" || mode == "production" {
//			cfg = zap.NewProductionConfig()
//		} else {
//			cfg = zap.NewDevelopmentConfig()
//		}
//
//		//cfg.OutputPaths = []string{
//		//	"./cheetah.log",
//		//}
//		zapLogger, err = cfg.Build()
//		if err != nil {
//			panic("init zap logger failed")
//		}
//	})
//
//	return zapLogger
//}

func Logger() *zap.Logger {
	if zapLogger == nil {
		initZapLogger()
	}
	return zapLogger
}

func initZapLogger() *zap.Logger {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("load env error: ", err.Error())
		return nil
	}
	mode := os.Getenv("SYSTEM_MODE")
	var encoder zapcore.Encoder
	if strings.ToLower(mode) == "prod" || strings.ToLower(mode) == "production" {
		// It's easy to understand the specific meaning of setting some basic log formats, and it's not difficult to understand the zap source code directly
		encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			CallerKey:    "file",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		})
	} else {
		// It's easy to understand the specific meaning of setting some basic log formats, and it's not difficult to understand the zap source code directly
		encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "ts",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			CallerKey:    "file",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendInt64(int64(d) / 1000000)
			},
		})
	}

	// Implement two interfaces to judge the log level (in fact, zapcore.*Level itself is the interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.ErrorLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	logFileName := os.Getenv("SYSTEM_LOG_FILE")
	logFileNameErr := os.Getenv("SYSTEM_LOG_FILE_ERR")
	// Get io.Writer Abstract getWriter() of info and warn log files and implement it below
	infoWriter := getWriter(logFileName)
	warnWriter := getWriter(logFileNameErr)

	// Finally, create a specific Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel),
	)

	zapLogger = zap.New(core, zap.AddCaller()) // You need to pass in zap.AddCaller() to display the file name and the number of lines of the log point. It's a bit small
	zapLogger.Info("init zap log")
	return zapLogger
}

func getWriter(filename string) io.Writer {
	// The actual file name generated by the Logger generating rotatelogs is demo.log.YYmmddHH
	// demo.log is a link to the latest log
	// Save the logs within 7 days, and split the logs every 1 hour (whole point)
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d-%H", // No go style anti human format
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
