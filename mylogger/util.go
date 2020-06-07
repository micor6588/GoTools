package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 反向解析
func getLogString(level LogLevel) (string, error) {
	switch level {
	case DEBUG:
		return "DEBUG", nil
	case INFO:
		return "INFO", nil
	case WARAING:
		return "WARAING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		err := errors.New("未知日志级别")
		return "UNKONW", err

	}
}

// getInfo  获取日志文件信息
func getInfo(n int) (funcName, fileName string, lineNum int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return funcName, fileName, line
}
