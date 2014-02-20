// Package xlog acts as a trivial extension of Go's log package.
// It provides functions to prefix the logged output with different
// log levels.
package xlog

import (
	"io"
	"log"
)

func SetOutput(w io.Writer) {
	log.SetOutput(w)
}

func Debug(v ...interface{}) {
	log.Print(append([]interface{}{"[DEBUG] "}, v...)...)
}

func Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] "+format, v...)
}

func Info(v ...interface{}) {
	log.Print(append([]interface{}{"[INFO]  "}, v...)...)
}

func Infof(format string, v ...interface{}) {
	log.Printf("[INFO]  "+format, v...)
}

func Warn(v ...interface{}) {
	log.Print(append([]interface{}{"[WARN]  "}, v...)...)
}

func Warnf(format string, v ...interface{}) {
	log.Printf("[WARN]  "+format, v...)
}

func Error(v ...interface{}) {
	log.Print(append([]interface{}{"[ERROR] "}, v...)...)
}

func Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] "+format, v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(append([]interface{}{"[FATAL] "}, v...)...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf("[FATAL] "+format, v...)
}

func Panic(v ...interface{}) {
	log.Panic(append([]interface{}{"[PANIC] "}, v...)...)
}

func Panicf(format string, v ...interface{}) {
	log.Panicf("[PANIC] "+format, v...)
}

func Request(v ...interface{}) {
	log.Print(append([]interface{}{"[REQ]   "}, v...)...)
}

func Requestf(format string, v ...interface{}) {
	log.Printf("[REQ]   "+format, v...)
}
