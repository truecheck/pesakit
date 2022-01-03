package log

import (
	"log/syslog"
	"syscall"
)

const Syslog = syslog.LOG_SYSLOG

type readfd int

func (r readfd) Read(buf []byte) (int, error) {
	return syscall.Read(int(r), buf)
}

type writefd int

func (w writefd) Write(buf []byte) (int, error) {
	return syscall.Write(int(w), buf)
}

const (
	StdIn  = readfd(0)
	StdOut = writefd(1)
	StdErr = writefd(2)
)
