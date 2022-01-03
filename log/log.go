package log

import (
	"io"
	"log"
)

func init() {
	log.SetPrefix("[PESAKIT] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

type Params struct {
	Output io.Writer
	Prefix string
}

type Option func(params *Params)

func WithOutput(output io.Writer) Option {
	return func(params *Params) {
		params.Output = output
	}
}

func WithPrefix(prefix string) Option {
	return func(params *Params) {
		params.Prefix = prefix
	}
}
func New(options ...Option) *log.Logger {
	params := &Params{
		Output: StdErr,
		Prefix: "[PESAKIT] ",
	}

	for _, option := range options {
		option(params)
	}

	return log.New(params.Output, params.Prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
