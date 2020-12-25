package common

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	bytes := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", bytes)
}