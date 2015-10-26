package log

import (
	"testing"
	"time"
	"fmt"
)

func Test_All(t *testing.T) {
	
	time := time.Now()
	
	fmt.Println(time.Location().String())
	
	fmt.Println(time.Unix())
	
}