package log

import (
	"fmt"
	"log"
	"log/syslog"
)

var logger = NewLogger(syslog.LOG_LOCAL0, log.LstdFlags|log.Lshortfile)

func Info(msg string, vals ...interface{}) {
	logger.Info(msg, vals...)
}

func Warning(msg string, vals ...interface{}) {
	logger.Warning(msg, vals...)
}

func Debug(msg string, vals ...interface{}) {
	logger.Debug(msg, vals...)
}

func Error(msg string, vals ...interface{}) {
	logger.Error(msg, vals...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func NewLogger(facility syslog.Priority, logFlags int) *Logger {
	l := new(Logger)
	var err error

	if l.debug, err = syslog.NewLogger(syslog.LOG_DEBUG|facility, logFlags); err != nil {
		log.Printf("LOGGER/DEBUG: %v", err)
	}

	if l.warn, err = syslog.NewLogger(syslog.LOG_WARNING|facility, logFlags); err != nil {
		log.Printf("LOGGER/WARNING: %v", err)
	}

	if l.info, err = syslog.NewLogger(syslog.LOG_INFO|facility, logFlags); err != nil {
		log.Printf("LOGGER/INFO: %v", err)
	}

	if l.err, err = syslog.NewLogger(syslog.LOG_ERR|facility, logFlags); err != nil {
		log.Print("LOGGER/ERROR: %v", err)
	}

	return l
}

type Logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
}

func (l *Logger) Info(msg string, vals ...interface{}) {
	l.info.Output(3, fmt.Sprintf(msg, vals...))
}

func (l *Logger) Warning(msg string, vals ...interface{}) {
	l.warn.Output(3, fmt.Sprintf(msg, vals...))
}

func (l *Logger) Debug(msg string, vals ...interface{}) {
	l.debug.Output(3, fmt.Sprintf(msg, vals...))
}

func (l *Logger) Error(msg string, vals ...interface{}) {
	l.err.Output(3, fmt.Sprintf(msg, vals...))
}

func (l *Logger) Write(p []byte) (int, error) {
	if err := l.info.Output(3, string(p)); err != nil {
		return 0, err
	} else {
		return len(p), nil
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.err.Print(v...)
	log.Fatal(v...)
}