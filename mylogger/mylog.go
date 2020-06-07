package mylogger

import (
	"errors"
	"strings"
)

// ParesLogLevel 解析日志级别
func ParesLogLevel(str string) (LogLevel, error) {
	// 将字符串全部转换成大写
	str = strings.ToLower(str)
	switch str {
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "waraing":
		return WARAING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil

	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err

	}

}

// NewLog 构造函数
func NewLog(levelStr string) Logging {
	level, err := ParesLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logging{
		LevelLog: level,
	}
}
