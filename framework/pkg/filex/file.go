package filex

import (
	"os"
	"path/filepath"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDirExist(dir string) bool {
	if dir == "" {
		return false
	}

	if dir == "." || dir == "./" {
		return true
	}

	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return info.IsDir()
}

func IsFileExist(filename string) bool {
	if filename == "" {
		return false
	}

	info, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return info.Mode().IsRegular()
}

func OpenFile(filename string, flag int, perm os.FileMode) (*os.File, error) {

	if flag&(os.O_CREATE|os.O_WRONLY|os.O_RDWR) != 0 {
		dir := filepath.Dir(filename)
		if !IsDirExist(dir) {
			err := os.MkdirAll(dir, perm)
			if err != nil {
				return nil, err
			}
		}
	}

	return os.OpenFile(filename, flag, perm)
}
