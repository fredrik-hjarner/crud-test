package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thoas/go-funk"
)

// ListKeys Return all keys
func ListKeys() []string {
	var slice []string

	walker := func(path string, info os.FileInfo, err error) error {
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
	}

	err := filepath.Walk("data-dir", walker)
	if err != nil {
		log.Println(err)
	}
	return slice
}

// ListKeysWithPrefix Return the keys, of all keys, that start with prefix.
func ListKeysWithPrefix(prefix string) []string {
	hasPrefix := func(key string) bool {
		return strings.HasPrefix(key, prefix)
	}

	return funk.FilterString(ListKeys(), hasPrefix)
}
