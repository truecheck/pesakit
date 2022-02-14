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

package mock

import (
	"context"
	"fmt"
	"github.com/pesakit/pesalib/tigo"
	"sync"
)

var _ tigo.Service = (*TigoServiceMock)(nil)

type TigoServiceMock struct {
	max     float64
	min     float64
	mux     *sync.RWMutex
	clients []string
	token   string
}

func validateAmount(min, max, amount float64) error {

	if amount < min || amount > max {
		return fmt.Errorf("amount validation failed, amount: %f, min: %f, max: %f", amount, min, max)
	}

	return nil
}

func (t *TigoServiceMock) Disburse(ctx context.Context, request tigo.DisburseRequest) (response tigo.DisburseResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *TigoServiceMock) Push(ctx context.Context, request tigo.PushRequest) (response tigo.PushResponse, err error) {
	err = validateAmount(t.min, t.max, request.Amount)
	if err != nil {
		response = tigo.PushResponse{
			ResponseCode:        tigo.FailureCode,
			ResponseStatus:      false,
			ResponseDescription: "Failed",
			ReferenceID:         request.ReferenceID,
			Message:             err.Error(),
		}
		return response, nil
	}

	response = tigo.PushResponse{
		ResponseCode:        tigo.SuccessCode,
		ResponseStatus:      true,
		ResponseDescription: "Success",
		ReferenceID:         request.ReferenceID,
	}

	return
}

func (t *TigoServiceMock) Token(ctx context.Context) (tigo.TokenResponse, error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	t.mux.RLock()
	defer t.mux.RUnlock()

	return tigo.TokenResponse{
		AccessToken: "",
		TokenType:   "",
		ExpiresIn:   0,
		Error:       "",
		ErrorDesc:   "",
	}, nil
}
