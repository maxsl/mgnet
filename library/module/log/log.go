package log

import (
	"bufio"
	"os"
	"time"
	"fmt"
	"sync"

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
	
	dir			string		// 日志目录
	baseName 	string		// 日志文件名
	reporting	int			// 错误级别
	
	logFile		bufio.Writer// 
	
	switcher	Switcher	// 文件切换器
	
	
	
	flushTick	time.Ticker	// flush间隔
	switchTick	time.Ticker	// 切换文件间隔
	
	syncRWMutexSwitch	sync.RWMutex	// 锁，切换文件时
	
}

func NewMgLog(dir string, baseName string, reporting int, switcher Switcher) (*MgLog) {

	return &MgLog{
		dir			:	dir,
		baseName	:	baseName,
		switcher 	: 	switcher,
		reporting 	:	reporting,
		
		flushTick	: 	time.NewTicker(5 * time.Second),
	}
	
}

// 周期性刷新日志
func (this *MgLog) runFlush() {
	go func() {
		for now := range this.flushTick.C {
			
		}
	}()
}

// 关闭日志
func (this *MgLog) Close() {
	this.flushTick.Stop()
}

// 日志刷新到磁盘
func (this *MgLog) flush() {
	this.logFile.Flush()
}


func (this *MgLog) switchFile() {
	this.syncRWMutexSwitch.Lock()
	defer this.syncRWMutexSwitch.Unlock()
	
	
}


func (this *MgLog) Info(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()
	
	// time := time.Now()
	// Log.Info(fmt.Sprintf("Info:%d", time.Unix()) + msg, data...)
	
	this.logFile.WriteString(fmt.Sprintf(msg, data ...))
}

func (this *MgLog) Warn(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

//	time := time.Now()
//	Log.Info(fmt.Sprintf("Warn:%d", time.Unix()) + msg, data...)

	this.logFile.WriteString(fmt.Sprintf(msg, data ...))
}

func (this *MgLog) Error(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

//	time := time.Now()
//	Log.Info(fmt.Sprintf("Error:%d", time.Unix()) + msg, data...)
	
	this.logFile.WriteString(fmt.Sprintf(msg, data ...))
}

func (this *MgLog) Debug(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

//	time := time.Now()
//	Log.Info(fmt.Sprintf("Debug:%d", time.Unix()) + msg, data...)
	
	this.logFile.WriteString(fmt.Sprintf(msg, data ...))
}





