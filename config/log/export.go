package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]any

// LogFunction For big messages, it can be more efficient to pass a function
// and only call it if the log level is actually enables rather than
// generating the log message and then checking if the level is enabled
type LogFunction func() []any

var (
	// std is the name of the standard logger in stdlib `log`
	std = logrus.New()
)

// GetOutput gets the standard logger output.
func GetOutput() io.Writer {
	return std.Out
}

// SetLevel sets the standard logger level.
func SetLevel(level string) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		WithError(err).Fatal("set level error")
	}
	std.SetLevel(lv)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return std.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return std.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value any) *logrus.Entry {
	return std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields Fields) *logrus.Entry {
	return std.WithFields(logrus.Fields(fields))
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return std.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...any) {
	std.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...any) {
	std.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...any) {
	std.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...any) {
	std.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...any) {
	std.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...any) {
	std.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...any) {
	std.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...any) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...any) {
	std.Fatal(args...)
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn LogFunction) {
	std.TraceFn(logrus.LogFunction(fn))
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn LogFunction) {
	std.DebugFn(logrus.LogFunction(fn))
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn LogFunction) {
	std.PrintFn(logrus.LogFunction(fn))
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn LogFunction) {
	std.InfoFn(logrus.LogFunction(fn))
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn LogFunction) {
	std.WarnFn(logrus.LogFunction(fn))
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn LogFunction) {
	std.WarningFn(logrus.LogFunction(fn))
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn LogFunction) {
	std.ErrorFn(logrus.LogFunction(fn))
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn LogFunction) {
	std.PanicFn(logrus.LogFunction(fn))
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn LogFunction) {
	std.FatalFn(logrus.LogFunction(fn))
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...any) {
	std.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...any) {
	std.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...any) {
	std.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...any) {
	std.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...any) {
	std.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...any) {
	std.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...any) {
	std.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...any) {
	std.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...any) {
	std.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...any) {
	std.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...any) {
	std.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...any) {
	std.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...any) {
	std.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...any) {
	std.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...any) {
	std.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...any) {
	std.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...any) {
	std.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...any) {
	std.Fatalln(args...)
}
