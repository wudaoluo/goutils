package sys

import "os"

// PathExists 判断目录是否存在   存在: true,不存在: false
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}