package log

import (
	"bufio"
//	"os"
	"time"
	"fmt"
)

var (
	Log * MgLog
)

// 初始化全局log
func InitLog(dir string, reporting int) {
	
}

// 在全局日志中输出信息
func Info(msg string, data ...interface{}) {
	time := time.Now()
	Log.Info(fmt.Sprintf("Info:%d", time.Unix()) + msg, data...)
}

// 在全局日志中输出警告信息
func Warn(msg string, data ...interface{}) {
	time := time.Now()
	Log.Info(fmt.Sprintf("Warn:%d", time.Unix()) + msg, data...)
}

// 在全局日志中输出错误信息
func Error(msg string, data ...interface{}) {
	time := time.Now()
	Log.Info(fmt.Sprintf("Error:%d", time.Unix()) + msg, data...)
}

// 在全局日志中输出调试信息
func Debug(msg string, data ...interface{}) {
	time := time.Now()
	Log.Info(fmt.Sprintf("Debug:%d", time.Unix()) + msg, data...)
}




type MgLog struct {
	logFile		bufio.Writer
	reporting	int			// 错误级别
	switcher	Switcher	// 文件切换器
}

func NewMgLog(dir string, baseName string, reporting int, switcher Switcher) (*MgLog) {

	return &MgLog{
		switcher : switcher,
		reporting : reporting,
	}
}

func (this *MgLog) Info(msg string, data ...interface{}) {
	
}

func (this *MgLog) Warn(msg string, data ...interface{}) {
	
}

func (this *MgLog) Error(msg string, data ...interface{}) {
	
}

func (this *MgLog) Debug(msg string, data ...interface{}) {
	
}





