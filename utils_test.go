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

package pesakit

import "testing"

func Test_copyToDir(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name     string
		args     args
		wantPath string
		wantErr  bool
	}{
		{
			name: "",
			args: args{
				src: "pesakit.env",
				dst: "/home/masengwa/.pesakit",
			},
			wantErr:  false,
			wantPath: "/home/masengwa/.pesakit/pesakit.env",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := copyToDir(tt.args.src, tt.args.dst)
			if (err != nil) != tt.wantErr {
				t.Errorf("copyToDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			if path != tt.wantPath {
				t.Errorf("copyToDir() path = %v, want %v", path, tt.wantPath)
			}
		})
	}
}
