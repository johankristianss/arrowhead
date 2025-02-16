package utils

import "os"

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func FileIsNewer(file1, file2 string) bool {
	info1, err1 := os.Stat(file1)
	info2, err2 := os.Stat(file2)
	if err1 != nil || err2 != nil {
		return false
	}
	return info1.ModTime().After(info2.ModTime())
}
