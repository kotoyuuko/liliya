package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// SetFormatter set log format
func SetFormatter(formatter logrus.Formatter) {
	log.SetFormatter(formatter)
}

// SetLevel set log level
func SetLevel(level logrus.Level) {
	log.SetLevel(level)
}

// SetOutput set output of log
func SetOutput(out io.Writer) {
	log.SetOutput(out)
}

// Trace writes something very low level to log
func Trace(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Trace(message)
	} else {
		log.Trace(message)
	}
}

// Debug writes useful debugging information to log
func Debug(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Debug(message)
	} else {
		log.Debug(message)
	}
}

// Info writes something noteworthy happened to log
func Info(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Info(message)
	} else {
		log.Info(message)
	}
}

// Warn writes messages you should probably take a look at to log
func Warn(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Warn(message)
	} else {
		log.Warn(message)
	}
}

// Error writes something failed to log
func Error(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Error(message)
	} else {
		log.Error(message)
	}
}

// Fatal writes something failed to log and calls os.Exit(1) after logging
func Fatal(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Fatal(message)
	} else {
		log.Fatal(message)
	}
}

// Panic writes something failed to log and calls panic() after logging
func Panic(message string, fields ...logrus.Fields) {
	if len(fields) > 0 {
		log.WithFields(fields[0]).Panic(message)
	} else {
		log.Panic(message)
	}
}