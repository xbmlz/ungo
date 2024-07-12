package log

// Debug log
func (z *zapLogger) Debug(v ...any) {
	if z.conf.Level <= LevelDebug {
		z.slog.Debug(v...)
	}
}

// Debugf log
func (z *zapLogger) Debugf(format string, v ...any) {
	if z.conf.Level <= LevelDebug {
		z.slog.Debugf(format, v...)
	}
}

// Info log
func (z *zapLogger) Info(v ...any) {
	if z.conf.Level <= LevelInfo {
		z.slog.Info(v...)
	}
}

// Infof log
func (z *zapLogger) Infof(format string, v ...any) {
	if z.conf.Level <= LevelInfo {
		z.slog.Infof(format, v...)
	}
}

// Warn log
func (z *zapLogger) Warn(v ...any) {
	if z.conf.Level <= LevelWarn {
		z.slog.Warn(v...)
	}
}

// Warnf log
func (z *zapLogger) Warnf(format string, v ...any) {
	if z.conf.Level <= LevelWarn {
		z.slog.Warnf(format, v...)
	}
}

// Error log
func (z *zapLogger) Error(v ...any) {
	if z.conf.Level <= LevelError {
		z.slog.Error(v...)
	}
}

// Errorf log
func (z *zapLogger) Errorf(format string, v ...any) {
	if z.conf.Level <= LevelError {
		z.slog.Errorf(format, v...)
	}
}
