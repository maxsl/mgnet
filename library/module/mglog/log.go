package mglog

import (
	"bufio"
	"os"
	"time"
	"fmt"
	"sync"
	"log"
	"runtime"
)

var (
	Log * MgLog
)

// 初始化全局log
func InitLog(dir string, baseName string, reporting int, switcher Switcher) {
	Log = NewMgLog(dir, baseName, reporting, switcher)
}

func Info(msg string, data ...interface{}) {
	Log.Info(msg, data ...)
}

func Warn(msg string, data ...interface{}) {
	Log.Info(msg, data ...)
}

func Error(msg string, data ...interface{}) {
	Log.Info(msg, data ...)
}

func Debug(msg string, data ...interface{}) {
	Log.Info(msg, data ...)
}




type MgLog struct {
	dir			string		// 日志目录
	baseName 	string		// 日志文件名
	reporting	int			// 错误级别

	logFile		*os.File	// 文件句柄
	logBufio	*bufio.Writer// 缓冲写入句柄
	logger		*log.Logger	// 日志句柄
	
	switcher	Switcher	// 文件切换器
	
	syncRWMutexSwitch	sync.RWMutex	// 锁，切换文件时

	flushDuration	time.Duration	// flush间隔
}

// 创建日志
func NewMgLog(dir string, baseName string, reporting int, switcher Switcher) (*MgLog) {

	mglog := &MgLog{
		dir			:	dir,
		baseName	:	baseName,
		switcher 	: 	switcher,
		reporting 	:	reporting,
		
		flushDuration: 	2 * time.Second,
	}
	
	go func() {
		for {
			mglog.switchFile()
			time.Sleep(mglog.switcher.NextDuration())
		}
	}()
	
	runtime.Gosched()
	
	go func() {
		for {
			mglog.flush()
			time.Sleep(mglog.flushDuration)
		}
	}()
	
	return mglog
}

// 关闭日志
func (this *MgLog) Close() {
	this.syncRWMutexSwitch.Lock()
	defer this.syncRWMutexSwitch.Unlock()	
	this.logBufio.Flush()	
	this.logFile.Close()
}

// 切换文件
func (this *MgLog) switchFile() {
	this.syncRWMutexSwitch.Lock()
	defer this.syncRWMutexSwitch.Unlock()

	if this.logFile != nil {
		this.logFile.Close()
	}
	
	this.logFile, _ = os.OpenFile(this.dir + "/" + this.switcher.Filename(this.baseName), os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0777)
	
	if this.logFile == nil {
		fmt.Println("logFile error")
	}
	
	this.logBufio = bufio.NewWriterSize(this.logFile, 1024000)
	
	if this.logBufio == nil {
		fmt.Println("logBufio error")
	}
	
	this.logger = log.New(this.logBufio, "", log.LstdFlags)
	
	if this.logger == nil {
		fmt.Println("logger error")
	}
}

// 日志刷新到磁盘
func (this *MgLog) flush() {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

	this.logBufio.Flush()
}

func (this *MgLog) Debug(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

	if this.reporting > E_DEBUG {
		return
	}

	if this.reporting == E_ALL {
		log.Printf("DEBUG : " + msg, data ...)
	}

	this.logger.Printf("DEBUG : " + msg, data ...)
}

func (this *MgLog) Info(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()
	
	if this.reporting > E_INFO {
		return
	}

	this.logger.Printf("INFO : " + msg, data ...)
}

func (this *MgLog) Warn(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

	if this.reporting > E_WARN {
		return
	}

	this.logger.Printf("WARN : " + msg, data ...)
}

func (this *MgLog) Error(msg string, data ...interface{}) {
	this.syncRWMutexSwitch.RLock()
	defer this.syncRWMutexSwitch.RUnlock()

	if this.reporting > E_ERROR {
		return
	}

	this.logger.Printf("ERROR : " + msg, data ...)
}