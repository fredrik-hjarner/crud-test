package utils

import (
	"math/rand"
	"strconv"
)

// CreateUUID ...
func CreateUUID() string {
	return strconv.Itoa(rand.Intn(1000000))
}
