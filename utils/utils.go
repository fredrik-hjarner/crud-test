package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/thoas/go-funk"
)

// ListKeys Return all keys
func ListKeys(namespace string) []string {
	var slice []string

	walker := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		key := strings.ReplaceAll(path, "\\", "/")
		key = strings.TrimPrefix(key, fmt.Sprintf("data-dir/%v", namespace))
		key = strings.TrimPrefix(key, "/")
		if key != "" && info.IsDir() == false {
			slice = append(slice, key)
		}
		return nil
	}

	err := filepath.Walk(fmt.Sprintf("data-dir/%v", namespace), walker)
	if err != nil {
		log.Println(err)
	}
	return slice
}

// ListKeysWithPrefix Return the keys, of all keys, that start with prefix.
func ListKeysWithPrefix(namespace string, prefix string) []string {
	hasPrefix := func(key string) bool {
		return strings.HasPrefix(key, prefix)
	}

	return funk.FilterString(ListKeys(namespace), hasPrefix)
}
