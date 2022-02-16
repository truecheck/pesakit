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

package cli

import (
	"fmt"
	"syscall"
)

type readfd int

func (r readfd) Read(buf []byte) (int, error) {
	n, err := syscall.Read(int(r), buf)
	if err != nil {
		return -1, fmt.Errorf("pesalib/io read error: %w", err)
	}

	return n, nil
}

type writefd int

func (w writefd) Write(buf []byte) (int, error) {
	n, err := syscall.Write(int(w), buf)
	if err != nil {
		return -1, fmt.Errorf("pesalib/io write error %w", err)
	}

	return n, nil
}

type readWriter int

func (rw readWriter) Read(buf []byte) (int, error) {
	n, err := syscall.Read(int(rw), buf)
	if err != nil {
		return -1, fmt.Errorf("pesalib/io read error: %w", err)
	}

	return n, nil
}

func (rw readWriter) Write(buf []byte) (int, error) {
	n, err := syscall.Write(int(rw), buf)
	if err != nil {
		return -1, fmt.Errorf("pesalib/io write error %w", err)
	}

	return n, nil
}

const (
	Stdin  = readfd(0)
	Stdout = writefd(1)
	Stderr = writefd(2)
	Stdio  = readWriter(1)
)
