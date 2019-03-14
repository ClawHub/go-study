package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MainLogger *zap.Logger
var GatewayLogger *zap.Logger

func init() {
	MainLogger = NewLogger("./logs/main.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
	GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
}
