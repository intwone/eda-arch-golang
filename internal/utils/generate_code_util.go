package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateCodeUtil() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	result := strconv.Itoa(code)
	return result
}
