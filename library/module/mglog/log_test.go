package mglog

import (
	"testing"
	"fmt"
)

func Test_All(t *testing.T) {

	mglog := NewMgLog("./", "wulei", 1, &SWITCHER_HOUR)
	
	mglog.Info("hello")
	mglog.Error("hello")
	mglog.Debug("hello")
	
	mglog.flush()
	fmt.Println(mglog)

	fmt.Println("a")
	
}