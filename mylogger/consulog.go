package mylogger

import (
	"fmt"
	"time"
)

// LogLevel 日志级别
type LogLevel uint16

const (
	DEBUG LogLevel = iota
	INFO
	WARAING
	ERROR
	FATAL
	UNKNOWN
)

// 往终端写入日志
// Logging 日志结构体
type Logging struct {
	LevelLog LogLevel
}

// enable 比较日志级别的大小
func (l Logging) enable(logLevel LogLevel) bool {
	return l.LevelLog <= logLevel
}

// logMsg记录日志信息
func (l Logging) logMsg(level LogLevel, format string, a ...interface{}) {
	if l.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNum := getInfo(3)
		str, err := getLogString(level)
		if err != nil {
			fmt.Println("反向解析日志内容失败")
		}
		fmt.Printf("[%s] [%s] [%s: %s: %d] %s,%s\n", now.Format("2006-01-02 15:04:05"), str, funcName, fileName, lineNum, format, msg)
	}
}

// Debug 指出细粒度信息事件对调试应用程序是非常有帮助的
func (l Logging) Debug(format string, a ...interface{}) {

	l.logMsg(DEBUG, format, a...)

}

// Info 表明消息在粗粒度级别上突出强调应用程序的运行过程
func (l Logging) Info(format string, a ...interface{}) {
	l.logMsg(INFO, format, a...)

}

// Warning 表明会出现潜在错误的情形
func (l Logging) Warning(format string, a ...interface{}) {
	l.logMsg(WARAING, format, a...)
}

// Error 指出虽然发生错误事件，但仍然不影响系统的继续运行
func (l Logging) Error(format string, a ...interface{}) {
	l.logMsg(ERROR, format, a...)

}

// Fatal 指出每个严重的错误事件将会导致应用程序的退出
func (l Logging) Fatal(format string, a ...interface{}) {
	l.logMsg(FATAL, format, a...)
}
