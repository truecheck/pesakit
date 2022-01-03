/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package env

import (
	"fmt"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
)

// LoadConfigFrom will read your env file(s) and load them into ENV for this process.
// Call this function as close as possible to the start of your program (ideally in main)
// If you call Load without any args it will default to loading .env in the current path
// You can otherwise tell it which files to load (there can be more than one) like
//
//		env.LoadConfigFrom("fileone", "filetwo")
//
// It's important to note that it WILL NOT OVERRIDE an env variable that already exists
//- consider the .env file to set dev vars or sensible defaults
func LoadConfigFrom(paths ...string) error {
	if err := dotenv.Load(paths...); err != nil {
		return fmt.Errorf("failed to load config files from %v .env file: %w", paths, err)
	}
	return nil
}

func get(key string, defaultValue interface{}) interface{} {
	var strValue string
	if strValue = os.Getenv(key); strValue == "" {
		return defaultValue
	}

	return strValue

}

func Set(key string, defaultValue interface{}) error {
	return os.Setenv(key, String(key, defaultValue))
}

func ReadSetString(key string, defaultValue string) (string, error) {
	err := Set(key, String(key, defaultValue))
	if err != nil {
		return "", fmt.Errorf("could not set the env %s, %w", key, err)
	}
	return String(key, defaultValue), nil
}

func ReadSetInt64(key string, defaultValue int64) (int64, error) {
	err := Set(key, String(key, defaultValue))
	if err != nil {
		return -1, fmt.Errorf("could not set the env %s, %w", key, err)
	}
	return Int64(key, defaultValue), nil
}

func ReadSetBool(key string, defaultValue bool) (bool, error) {
	err := Set(key, String(key, defaultValue))
	if err != nil {
		return false, fmt.Errorf("could not set the env %s, %w", key, err)
	}
	return Bool(key, defaultValue), nil
}

func ReadSetFloat64(key string, defaultValue float64) (float64, error) {
	err := Set(key, String(key, defaultValue))
	if err != nil {
		return -1, fmt.Errorf("could not set the env %s, %w", key, err)
	}
	return Float64(key, defaultValue), nil
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
