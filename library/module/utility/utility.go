package utility

import (
	"fmt"
	"crypto/sha1"
	"time"
	"strconv"
)

// 生成 Session
func GenSessionId(ptId int32, serverId int32) string {
	now := time.Now()

	timestamp := strconv.FormatInt(now.UnixNano(), 36)

	id := fmt.Sprintf("%x", sha1.Sum([]byte(timestamp)))

	return id
}

