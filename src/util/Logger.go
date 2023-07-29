package util

import (
	"log"
	"os"
	"strconv"
	"time"
)

const (
	Debug logLevel = iota
	Info
	Warn
	Error
)

// logLevel	日志等级
type logLevel int

// f 	异常处理函数
var f func(r any)

// flag	是否开启debug模式
var flag bool

// LoggerInit
// @Description: 初始化日志
// @param        f 异常处理函数
// @param        F 是否开启debug模式
func LoggerInit(f func(r any), F logLevel) {
	switch F {
	case Debug:
		flag = true
	case Info, Warn, Error:
		flag = false
	}
	setExceptionFunc(f)
	if !exists("./log") {
		_ = os.Mkdir("./log", 0644)
	}
}

// Loglevel
// @Description: 日志等级
// @param        level   日志等级
// @param        name    日志名称
// @param        message 日志内容
func Loglevel(level logLevel, name string, message string) {
	var logger *log.Logger
	defer func() {
		r := recover()
		if r != nil {
			f(r)
		}
	}()

	// 创建日志对象
	switch level {
	case Debug:
		logger = log.New(os.Stdout, name+" == "+" ["+"debug"+"] ", log.Ldate|log.Ltime)
	case Info:
		logger = log.New(os.Stdout, name+" == "+" ["+"info"+"] ", log.Ldate|log.Ltime)
	case Warn:
		logger = log.New(os.Stdout, name+" == "+" ["+"warn"+"] ", log.Ldate|log.Ltime)
	case Error:
		logger = log.New(os.Stdout, name+" == "+" ["+"error"+"] ", log.Ldate|log.Ltime)
	}

	// 记录日志
	switch level {
	case Debug:
		if flag {
			logger.Println(message)
			recordFile(message, level, logger)
		}
	case Info, Warn, Error:
		logger.Println(message)
		recordFile(message, level, logger)
	default:
		log.Panic("无此选项")
	}
}

// setExceptionFunc
// @Description: 设置异常处理函数
// @param        exceptionFunc 异常处理函数
func setExceptionFunc(exceptionFunc func(r any)) {
	f = exceptionFunc
}

// recordFile
// @Description: 记录日志到文件
// @param        message 日志内容
// @param        level   日志等级
// @param        logger  日志对象
func recordFile(message string, level logLevel, logger *log.Logger) {
	var FileLevel string
	switch level {
	case Debug:
		FileLevel = "debug"
	case Info:
		FileLevel = "Info"
	case Warn:
		FileLevel = "Warn"
	case Error:
		FileLevel = "Error"
	}

	// 获取当前时间，创建当日文件夹
	year, month, day := time.Now().Date()
	t := strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day)
	filename := "./log/" + t
	if !exists(filename) {
		_ = os.Mkdir(filename, 0644)
	}

	// 创建当日日志文件
	file, err := os.OpenFile(filename+"/"+FileLevel+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		f("打开日志文件异常")
	}

	// 设置日志输出到文件
	logger.SetOutput(file)

	// 记录日志
	logger.Println(message)
	err = file.Close()
	if err != nil {
		f("关闭日志文件异常")
		return
	}
}

// exists
// @Description: 判断文件是否存在
// @param        path 文件路径
// @return       bool 是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
