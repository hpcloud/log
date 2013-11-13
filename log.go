package log

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

type Logger struct {
	*log.Logger
}

var std = &Logger{log.New(os.Stderr, "", log.LstdFlags)}

func Info(v ...interface{}) {
	std.Output(2, "INFO -- "+fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf("INFO -- "+format, v...))
}

func Warn(v ...interface{}) {
	std.Output(2, "WARN -- "+fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf("WARN -- "+format, v...))
}

func Error(v ...interface{}) {
	std.Output(2, "ERROR -- "+fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf("ERROR -- "+format, v...))
}

func Fatal(v ...interface{}) {
	std.Output(2, "FATAL -- "+fmt.Sprint(v...)+"\n"+string(debug.Stack()))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf("FATAL -- "+format, v...)+"\n"+string(debug.Stack()))
	os.Exit(1)
}

func Fatal0(format string, v ...interface{}) {
	std.Output(2, fmt.Sprintf("FATAL -- "+format, v...)+"\n"+string(debug.Stack()))
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(2, "INFO -- "+fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf("INFO -- "+format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(2, "WARN -- "+fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf("WARN -- "+format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(2, "ERROR -- "+fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf("ERROR -- "+format, v...))
}

// Fatal behaves mostly like log.Fatal but it also prints the current
// goroutine's stacktrace.
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(2, "FATAL -- "+fmt.Sprint(v...)+"\n"+string(debug.Stack()))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf("FATAL -- "+format, v...)+"\n"+string(debug.Stack()))
	os.Exit(1)
}

// Fatal0 is like Fatalf, but does not exit (useful to performing cleanup
// routines after displaying the fatal message)
func (l *Logger) Fatal0(format string, v ...interface{}) {
	l.Output(2, fmt.Sprintf("FATAL -- "+format, v...)+"\n"+string(debug.Stack()))
}

func New() *Logger {
	return &Logger{log.New(os.Stderr, "", log.LstdFlags)}
}
