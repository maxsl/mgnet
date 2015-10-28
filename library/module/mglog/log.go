package mglog

import (
	"bufio"
	"os"
	"time"
	"fmt"
	"sync"
	"log"
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

	logFile		*os.File	// 文件句柄
	logBufio	*bufio.Writer// 缓冲写入句柄
	logger		*log.Logger	// 日志句柄
	
	switcher	Switcher	// 文件切换器
	
	flushTick	*time.Ticker	// flush间隔
	switchTick	*time.Ticker	// 切换文件间隔
	
	syncRWMutexSwitch	sync.RWMutex	// 锁，切换文件时

	syncGroupWait		sync.WaitGroup	// 等待锁，退出等待
}

func NewMgLog(dir string, baseName string, reporting int, switcher Switcher) (*MgLog) {

	mglog := &MgLog{
		dir			:	dir,
		baseName	:	baseName,
		switcher 	: 	switcher,
		reporting 	:	reporting,
		
		flushTick	: 	time.NewTicker(2 * time.Second),
	}
	
	mglog.switchFile()
	
	return mglog

}

//// 周期性刷新日志
//func (this *MgLog) run() {
//	go func() {
//		// Flush
//		this.syncGroupWait.add(1)
//		for now := range this.flushTick.C {
//			this.flush()
//		}
//		this.syncGroupWait.Done()
//	}()
//}



// 关闭日志
func (this *MgLog) Close() {
	this.flushTick.Stop()
}

// 日志刷新到磁盘
func (this *MgLog) flush() {
	this.logBufio.Flush()
}

// 切换文件
func (this *MgLog) switchFile() {
	this.syncRWMutexSwitch.Lock()
	defer this.syncRWMutexSwitch.Unlock()
	
	this.logFile, _ = os.OpenFile(this.dir + "/" + this.switcher.Filename(this.baseName) , os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)
	
	this.logBufio = bufio.NewWriter(this.logFile)
	
	this.logger = log.New(this.logBufio, "", log.LstdFlags)
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
