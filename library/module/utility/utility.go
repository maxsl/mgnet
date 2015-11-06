package utility

import (
	"fmt"
	"crypto/sha1"
	"time"
)

// 生成 Session
func GenSessionId(ptId uint32, serverId uint32, slat string) (sessId string) {
	now := time.Now()

	oriId := fmt.Sprintf("%d%d%s%d", ptId, serverId, slat, now.UnixNano())

	sessId = fmt.Sprintf("%x", sha1.Sum([]byte(oriId)))
	
	return sessId
}

