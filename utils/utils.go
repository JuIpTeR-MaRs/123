package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

// MD5String 返回字符串的 MD5 哈希（小写十六进制）
func MD5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// n 日前开始时间 2006-01-02 00:00:00
func GetStartTime(n int) string {
	t := time.Now()
	t = t.AddDate(0, 0, -n)
	return t.Format("2006-01-02") + " 00:00:00"
}

// n 日前结束时间 2006-01-02 23:59:59
func GetEndTime(n int) string {
	t := time.Now()
	t = t.AddDate(0, 0, -n)
	return t.Format("2006-01-02") + " 23:59:59"
}
