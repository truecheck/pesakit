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

package errors

var _ error = (*PesakitError)(nil)

type PesakitError struct {
	code        int
	description string
	message     string
	err         error
}

func New(code int, message string, description string, err error) *PesakitError {
	return &PesakitError{
		code:        code,
		description: description,
		message:     message,
		err:         err,
	}
}

func (e *PesakitError) Code() int {
	return e.code
}

func (e *PesakitError) Unwrap() error { return e.err }

func (e *PesakitError) Description() string {
	return e.description
}

func (e *PesakitError) Message() string {
	return e.message
}

func (e *PesakitError) Err() error {
	return e.err
}

func (e *PesakitError) Error() string {
	if e.err != nil {
		return e.message + ": " + e.err.Error()
	}
	return e.message
}
