package diskv

import (
	"strings"

	"github.com/peterbourgon/diskv"
)

func keyTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last],
	}
}

// If you provide an AdvancedTransform, you must also provide its
// inverse:
func keyInverseTransform(pathKey *diskv.PathKey) (key string) {
	txt := pathKey.FileName[len(pathKey.FileName):]
	if txt != ".txt" {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)]
}

// Diskv Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
var Diskv = diskv.New(diskv.Options{
	BasePath:          "data-dir",
	AdvancedTransform: keyTransform,
	InverseTransform:  keyInverseTransform,
	CacheSizeMax:      1024 * 1024,
})
