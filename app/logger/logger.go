package logger

import (
	"archilltect-sigma/app/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

func Init() (err error) {
	//writer := getLogWriter()
	//encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(settings.Config.LogConfig.Level))
	if err != nil {
		return
	}
	core := zapcore.NewTee(
		//zapcore.NewCore(encoder, writer, l),
		zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)
	logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   settings.Config.LogConfig.FileName,
		MaxSize:    settings.Config.LogConfig.MaxSize,
		MaxBackups: settings.Config.LogConfig.MaxBackups,
		MaxAge:     settings.Config.LogConfig.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
