package unlog

import "time"

// DefaultLogger is default logger.
var DefaultLogger = NewZapLogger(Config{
	Level:        LevelInfo,
	Path:         "logs",
	Name:         "app.log",
	MaxAge:       7 * 24 * time.Hour,
	RotationTime: 24 * time.Hour,
})

// Logger logger interface
type Logger interface {
	Debug(v ...any)
	Debugf(format string, v ...any)
	Info(v ...any)
	Infof(format string, v ...any)
	Warn(v ...any)
	Warnf(format string, v ...any)
	Error(v ...any)
	Errorf(format string, v ...any)
}
