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

import (
	"context"
	"errors"
	"fmt"
	"github.com/pesakit/pesakit/env"
	"github.com/pesakit/pesakit/flags"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	errServerNotinitialized = errors.New("server not initialized")
	serverClosingTimeout    = 5 * time.Second
)

const (
	flagCallbackHost       = "host"
	flagCallbackPort       = "port"
	flagCallbackPath       = "path"
	flagCallbackOperation  = "operation"
	usageCallbackHost      = "callback host"
	usageCallbackPort      = "callback port"
	usageCallbackPath      = "callback path"
	usageCallbackOperation = "callback operation"
	envCallbackHost        = "PESAKIT_CALLBACK_HOST"
	envCallbackPort        = "PESAKIT_CALLBACK_PORT"
	envCallbackPath        = "PESAKIT_CALLBACK_PATH"
	envCallbackOps         = "PESAKIT_CALLBACK_OPERATION"
	defaultCallbackHost    = "localhost"
	defaultCallbackPort    = 8080
	defaultCallbackPath    = "/callbacks"
	defaultCallbackOps     = "push"
)

const (
	push callbackOperation = "push"
)

type callbackServer struct {
	mu     *sync.RWMutex
	server *http.Server
	host   string
	port   int64
	path   string
	op     callbackOperation
	mno    mno.Mno
	logger *log.Logger
	writer io.Writer
}

func newCallbackServer(params *callbackParams, logger *log.Logger, writer io.Writer) *callbackServer {
	srv := &callbackServer{
		mu:     &sync.RWMutex{},
		host:   params.Host,
		port:   params.Port,
		path:   params.Path,
		op:     params.Operation,
		mno:    params.Mno,
		logger: logger,
		writer: writer,
	}

	if params.Mno == mno.Airtel {
		srv.makeServer(srv.genericHandlerFunc)
	} else {
		srv.makeServer(nil)
	}
	return srv
}

func (s *callbackServer) get() *http.Server {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.server
}

func (s *callbackServer) genericHandlerFunc(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
	// dump the whole request
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		s.logger.Printf("error dumping request: %s", err)
	}
	_, err = s.writer.Write(dump)
	if err != nil {
		s.logger.Printf("error writing request: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		s.logger.Printf("error writing response: %s", err)
		return
	}
}

// serveMux is the HTTP multiplexer used by the callback server.
func (s *callbackServer) serveMux(handler http.HandlerFunc) *http.ServeMux {
	if handler == nil {
		handler = s.genericHandlerFunc
	}
	mux := http.NewServeMux()
	mux.HandleFunc(s.path, handler)
	return mux
}

func (s *callbackServer) makeServer(handler http.HandlerFunc) {
	mux := s.serveMux(handler)
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", s.host, s.port),
		Handler:           mux,
		ReadTimeout:       20 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       20 * time.Second,
		MaxHeaderBytes:    1 << 20,
		ErrorLog:          s.logger,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.server = srv
}

func (s *callbackServer) ListenAndServe() error {
	s.logger.Printf("Starting callback server on %s:%d", s.host, s.port)
	srv := s.get()
	if srv == nil {
		return errServerNotinitialized
	}

	return srv.ListenAndServe()
}

func (s *callbackServer) Shutdown() error {
	s.logger.Printf("Shutting down callback server on %s:%d", s.host, s.port)
	srv := s.get()
	if srv == nil {
		return errServerNotinitialized
	}
	ctx, cancel := context.WithTimeout(context.Background(), serverClosingTimeout)
	defer cancel()

	return srv.Shutdown(ctx)
}

type callbackOperation string

type callbackParams struct {
	Mno       mno.Mno
	Host      string
	Port      int64
	Path      string
	Operation callbackOperation
}

func setCallbackFlags(cmd *cobra.Command) {
	var (
		callbackHost = env.String(envCallbackHost, defaultCallbackHost)
		callbackPort = env.Int64(envCallbackPort, defaultCallbackPort)
		callbackPath = env.String(envCallbackPath, defaultCallbackPath)
		callbackOps  = env.String(envCallbackOps, defaultCallbackOps)
	)

	flags.SetMnoFlag(cmd, flags.PERSISTENT)
	strVar := cmd.PersistentFlags().StringVar
	intVar := cmd.PersistentFlags().Int64Var
	strVar(&callbackHost, flagCallbackHost, callbackHost, usageCallbackHost)
	intVar(&callbackPort, flagCallbackPort, callbackPort, usageCallbackPort)
	strVar(&callbackPath, flagCallbackPath, callbackPath, usageCallbackPath)
	strVar(&callbackOps, flagCallbackOperation, callbackOps, usageCallbackOperation)
}

func loadCallbackParams(cmd *cobra.Command) (*callbackParams, error) {
	str := cmd.Flags().GetString
	host, err := str(flagCallbackHost)
	if err != nil {
		return nil, err
	}
	port, err := cmd.Flags().GetInt64(flagCallbackPort)
	if err != nil {
		return nil, err
	}
	path, err := str(flagCallbackPath)
	if err != nil {
		return nil, err
	}
	operation, err := str(flagCallbackOperation)
	if err != nil {
		return nil, err
	}
	mnoValue, err := flags.LoadMnoConfig(cmd, flags.PERSISTENT)
	if err != nil {
		return nil, err
	}

	return &callbackParams{
		Mno:       mnoValue,
		Host:      host,
		Port:      port,
		Path:      path,
		Operation: callbackOperation(operation),
	}, nil

}

func (app *App) callbacksCommand() {

	// callbacksCmd represents the callbacks command
	var callbacksCmd = &cobra.Command{
		Use:     "callbacks",
		Short:   "listen and logs http callbacks requests",
		Example: "pesakit callbacks --mno=tigo --port=9095 --path=/callbacks/tigo",
		Long: `listen and logs http callback requests from mobile money providers.
A new http server will be set up to listen to tigo callbacks requests.
at http://host:port/path/
`,
		Run: func(cmd *cobra.Command, args []string) {
			params, err := loadCallbackParams(cmd)
			if err != nil {
				app.Logger().Printf("error: %s\n", err)

				return
			}
			logger := app.Logger()
			writer := app.getWriter()
			server := newCallbackServer(params, logger, writer)
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
			go func() {
				<-c
				err := server.Shutdown()
				if err != nil {
					return
				}
				os.Exit(0)
			}()
			err = server.ListenAndServe()
			if err != nil {
				app.Logger().Printf("error: %v\n", err)

				return
			}
		},
	}
	setCallbackFlags(callbacksCmd)
	app.root.AddCommand(callbacksCmd)
}
