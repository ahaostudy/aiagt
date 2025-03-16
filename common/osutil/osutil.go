package osutil

import (
	"github.com/pkg/errors"
	"os"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}

func CopyFile(src, dst string) error {
	file, err := os.ReadFile(src)
	if err != nil {
		return errors.Wrap(err, "read source file error")
	}

	err = os.WriteFile(dst, file, 0644)
	if err != nil {
		return errors.Wrap(err, "write to dest file error")
	}

	return nil
}
