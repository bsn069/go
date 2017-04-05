package utils

import (
	"os"
)

func IsRootUser() bool {
	return os.Geteuid() != 0
}

func Env(strName string) string {
	return os.Getenv(strName)
}
