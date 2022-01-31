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

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFile copies a file from src to dst. If src or dst files do not exist
// it returns an error. src is expected to be a regular file and dst is expected
// to be a directory or else error is returned.
// copyFile if successful returns the new path of the file in dst.
func copyToDir(src, dst string) (string, error) {
	srcFileStat, err := os.Stat(src)
	if err != nil {
		return "", err
	}
	if !srcFileStat.Mode().IsRegular() {
		return "", fmt.Errorf("%s is not a regular file", src)
	}

	dstDirStat, err := os.Stat(dst)
	if err != nil {
		return "", err
	}
	if !dstDirStat.Mode().IsDir() {
		return "", fmt.Errorf("%s is not a directory", dst)
	}

	dstFile := filepath.Join(dst, srcFileStat.Name())
	srcFile, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()
	destination, err := os.Create(dstFile)
	if err != nil {
		return "", err
	}
	defer destination.Close()
	_, err = io.Copy(destination, srcFile)
	if err != nil {
		return "", err
	}
	return dstFile, nil
}
