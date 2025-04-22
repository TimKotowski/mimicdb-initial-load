package internal

import (
	"os"
	"strconv"
)

func GetEnvString(value string) string {
	val, ok := os.LookupEnv(value)
	if !ok {
		return ""
	}

	return val
}

func GetEnvInt(value string) int64 {
	val, ok := os.LookupEnv(value)
	if !ok {
		return 0
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}

	return i
}
