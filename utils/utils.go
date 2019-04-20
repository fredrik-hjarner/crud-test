package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

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

func ListKeysWithPrefix(prefix string) []string {
	return Filter(ListKeys(), func(v string) bool {
		return strings.HasPrefix(v, prefix)
	})
}
