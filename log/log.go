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
