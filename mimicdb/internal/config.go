package internal

import (
	"os"
)

func GetTargetDBHost() string {
	return GetEnvString("TARGET_DB_NAME")
}

func GetTargetDBUser() string {
	return GetEnvString("TARGET_DB_USER")
}

func GetTargetDBName() string {
	return GetEnvString("TARGET_DB_USER")
}

func GetTargetDBPort() int64 {
	return GetEnvInt("TARGET_DB_PORT")
}

func GetSourceDBHost() string {
	val, ok := os.LookupEnv("SOURCE_DB_HOST")
	if !ok {
		return ""
	}

	return val
}

func GetSourceDBPort() string {
	val, ok := os.LookupEnv("SOURCE_DB_PORT")
	if !ok {
		return ""
	}

	return val
}

func GetSourceDBName() string {
	val, ok := os.LookupEnv("SOURCE_DB_NAME")
	if !ok {
		return ""
	}

	return val
}
