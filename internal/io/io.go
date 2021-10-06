package io

// https://dave.cheney.net/2017/06/11/go-without-package-scoped-variables
import (
	"syscall"
)

type readfd int

func (r readfd) Read(buf []byte) (int, error) {
	return syscall.Read(int(r), buf)
}

type writefd int

func (w writefd) Write(buf []byte) (int, error) {
	return syscall.Write(int(w), buf)
}

const (
	Stdin  = readfd(0)
	Stdout = writefd(1)
	Stderr = writefd(2)
)
