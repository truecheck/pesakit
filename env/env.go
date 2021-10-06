package env

import (
	"fmt"
	"os"
	"strconv"
)

func get(key string, defaultValue interface{}) interface{} {
	var strValue string
	if strValue = os.Getenv(key); strValue == "" {
		return defaultValue
	}

	return strValue

}

func String(key string, defaultValue interface{}) string {
	i := get(key, defaultValue)
	return fmt.Sprintf("%v", i)
}

func Bool(key string, defaultValue bool) bool {
	i := String(key, defaultValue)
	parseBool, _ := strconv.ParseBool(i)
	return parseBool
}

func Int64(key string, defaultValue int64) int64 {
	i := String(key, defaultValue)
	parseInt, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		text := fmt.Sprintf("environment variable %s is suppored to be of type int64 but it reads as ${%s}, %s will be=%v", key, key, key, defaultValue)
		textEnv := os.ExpandEnv(text)
		_, _ = fmt.Fprintln(os.Stderr, textEnv)
	}
	return parseInt
}

func Float64(key string, defaultValue float64) float64 {
	i := String(key, defaultValue)
	parseFloat, err := strconv.ParseFloat(i, 64)
	if err != nil {
		text := fmt.Sprintf("environment variable %s is suppored to be of type float64 but it reads as ${%s}, %s will be=%v", key, key, key, defaultValue)
		textEnv := os.ExpandEnv(text)
		_, _ = fmt.Fprintln(os.Stderr, textEnv)
	}
	return parseFloat
}
