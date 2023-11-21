package pkg

import "os"

// PathExists 检查文件或目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 文件或目录不存在
		return false, nil
	}
	// 发生其他错误
	return false, err
}
