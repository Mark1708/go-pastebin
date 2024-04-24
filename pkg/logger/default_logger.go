package logger

import (
	defLog "log"
	"os"

	"github.com/Mark1708/go-pastebin/pkg/color"
)

type defaultLogger struct {
	infoLog  *defLog.Logger
	debugLog *defLog.Logger
	warnLog  *defLog.Logger
	errorLog *defLog.Logger
}

func NewDefaultLogger() Logger {
	flag := defLog.Ldate | defLog.Ltime | defLog.Lmicroseconds | defLog.LUTC | defLog.Lshortfile
	return &defaultLogger{
		infoLog:  defLog.New(os.Stdout, "INFO\t", flag),
		debugLog: defLog.New(os.Stdout, "DEBUG\t", flag),
		warnLog:  defLog.New(os.Stdout, "WARN\t", flag),
		errorLog: defLog.New(os.Stdout, "ERROR\t", flag),
	}
}

func (z *defaultLogger) Debug(msg string) {
	z.debugLog.Print(color.White + msg + color.Reset)
}

func (z *defaultLogger) Debugf(template string, args ...interface{}) {
	z.debugLog.Printf(color.White+template+color.Reset, args...)
}

func (z *defaultLogger) Info(msg string) {
	z.infoLog.Print(color.Green + msg + color.Reset)
}

func (z *defaultLogger) Infof(template string, args ...interface{}) {
	z.infoLog.Printf(color.Green+template+color.Reset, args...)
}

func (z *defaultLogger) Error(msg string) {
	z.errorLog.Print(color.Red + msg + color.Reset)
}

func (z *defaultLogger) Errorf(template string, args ...interface{}) {
	z.errorLog.Printf(color.Red+template+color.Reset, args...)
}

func (z *defaultLogger) Warn(msg string) {
	z.warnLog.Printf(color.Yellow + msg + color.Reset)
}

func (z *defaultLogger) Warnf(template string, args ...interface{}) {
	z.warnLog.Printf(color.Yellow+template+color.Reset, args...)
}

func (z *defaultLogger) Fatal(msg string) {
	z.errorLog.Fatal(msg)
}

func (z *defaultLogger) Fatalf(template string, args ...interface{}) {
	z.errorLog.Fatalf(template, args...)
}

func (z *defaultLogger) Panic(msg string) {
	z.errorLog.Panic(msg)
}

func (z *defaultLogger) Panicf(template string, args ...interface{}) {
	z.errorLog.Panicf(template, args...)
}

func (z *defaultLogger) With(_ ...interface{}) Logger {
	z.Error("method with not implemented for default logger")
	return z
}
