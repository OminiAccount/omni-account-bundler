package utils

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, fmt.Errorf("mkdir failed![%v]\n", err)
		} else {
			return true, nil
		}
	}
	return false, err
}
