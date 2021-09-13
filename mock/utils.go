package mock

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/google/uuid"
)

// NewUUID 生成uuid
func NewUUID() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

//NewMD5 生成md5
func NewMD5(str string) string {
	if len(str) == 0 {
		return ""
	}
	sign := md5.New()
	sign.Write([]byte(str))
	strSign := sign.Sum(nil)
	strMD5 := hex.EncodeToString(strSign)
	return strMD5
}

//NewUnixtime 获取当前的unixtime
func NewUnixtime() uint {
	return uint(time.Now().Unix())
}
