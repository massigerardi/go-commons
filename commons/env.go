package commons

import (
	"os"
	"strconv"
	"strings"
)

func GetString(name string, defaultValue string) string {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}

func GetList(name string, defaultValue []string) []string {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	return strings.Split(value, ",")
}

func GetBool(name string, defaultValue bool) bool {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	return strings.ToLower(value) == "true"
}

func GetInt(name string, defaultValue int64) int64 {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func GetFloat(name string, defaultValue float64) float64 {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func GetFloat32(name string, defaultValue float32) float32 {
	value, exists := os.LookupEnv(name)
	if !exists || value == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return defaultValue
	}
	return float32(intValue)
}
