package utils

import "os"

func GetEnvDefault(varName string, defaultVal string) string {
	value := os.Getenv(varName)
	if value == "" {
		return defaultVal
	}
	return value
}
