package mglog

import (
	"time"
	"fmt"
)

var (
	SWITCHER_HOUR 	SwitcherHour
	SWITCHER_DAY	SwitcherDay
)

type Switcher interface {
	NextDuration() (time.Duration)
	Filename(basename string) (string)
}

type SwitcherHour struct {
	
}

// 当前时间到下一小时间隔
func (this *SwitcherHour) NextDuration() (time.Duration) {
	now := time.Now()

	nextTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())

	dur := nextTime.Sub(now)
	
	return dur
}

// 当前文件名
func (this *SwitcherHour) Filename(basename string) (string) {
	now := time.Now()

	return fmt.Sprintf("%s_%d%d%d%d.log", basename, now.Year(), now.Month(), now.Day(), now.Hour())
}

type SwitcherDay struct {
	
}

// 当前时间到下一天间隔
func (this *SwitcherDay) NextDuration() (time.Duration) {
	now := time.Now()

	nextTime := time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour(), 0, 0, 0, now.Location())

	dur := nextTime.Sub(now)
	
	return dur
}

// 当前文件名
func (this *SwitcherDay) Filename(basename string) (string){
	now := time.Now()

	return fmt.Sprintf("%s_%d%d%d.log", basename, now.Year(), now.Month(), now.Day())
}