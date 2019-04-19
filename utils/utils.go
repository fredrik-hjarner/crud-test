package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListKeys() []string {
	var slice []string

	err := filepath.Walk("data-dir",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			key := strings.ReplaceAll(path, "\\", "/")
			key = strings.TrimPrefix(key, "data-dir")
			key = strings.TrimPrefix(key, "/")
			if key != "" && info.IsDir() == false {
				slice = append(slice, key)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return slice
}
