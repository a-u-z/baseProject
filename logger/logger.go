package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// fatal 直接壞掉
// panic 可以執行 defer

var L *zap.Logger

func Init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder // 設定時間 "2023-06-12T03:13:46.240Z"
	fileEncoder := zapcore.NewJSONEncoder(config)  // json 化
	// consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("log.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
	)
	L = zap.New(core, zap.AddCaller())
	// L = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
