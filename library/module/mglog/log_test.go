package mglog

import (
	"testing"
	"time"
	"runtime"
)

func Test_All(t *testing.T) {

	runtime.SetCPUProfileRate(4)

	mglog := NewMgLog("./", "wulei", 0, &SWITCHER_HOUR)
	
	go func () {		
		for {
			
			mglog.Debug(time.Now().String());
			
			time.Sleep(1 * time.Second)
		}
		
	}()
	
	
	time.Sleep(4 * time.Minute)
	
}