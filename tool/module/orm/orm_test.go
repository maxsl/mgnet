package orm

import (
	"testing"
	
)

var TestConfig []OrmConfig = []OrmConfig{
		// 用户表
		{
			"user",
			[]string{"UserId", "PtId", "Name", "Desc"},
			[]string{"int64", "string", "string", "string"},
			[]int{0},
			map[int]int{3:1024},
			map[int][]OrmJson{
				3 : []OrmJson{
						{"Name", "string"}, 
						{"Age", "int"} }}}}



func Test_FlushOrm(t *testing.T) {
	FlushOrm("/root/workspace/go/src/github.com/goodkele/mgnet/library/orm")
}
