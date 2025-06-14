package utils

import (
	"log"
	"path"
	"runtime"
)

func LogBusiness(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	fileName := "unknown"
	funcName := "unknown"
	if ok {
		fileName = path.Base(file)
		funcName = runtime.FuncForPC(pc).Name()
	}
	log.Printf("[BIZ] %s | %s:%d | %s", msg, fileName, line, funcName)
}
