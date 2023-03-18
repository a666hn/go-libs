package http_servers

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type iHttpServer interface {
	ListenAndServe() error
	SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}

// type HttpServer struct {
// 	IHttpServer
// 	Addr   string
// 	Logger *logrus.Logger
// }

type IHTTPInstanceServer interface {
	Start(quit chan os.Signal)
}

type httpServer struct {
	iHttpServer
	HttpContext context.Context
	Addr        string
	Logger      *logrus.Logger
}

type httpServerOptions func(o *httpServer)

func SetHttpContext(ctx context.Context) func(o *httpServer) {
	return func(o *httpServer) {
		o.HttpContext = ctx
	}
}

func SetHttpAddress(port string) func(o *httpServer) {
	return func(o *httpServer) {
		o.Addr = fmt.Sprintf("0.0.0.0:%s", port)
	}
}

func SetHttpLogger(logger *logrus.Logger) func(o *httpServer) {
	return func(o *httpServer) {
		o.Logger = logger
	}
}

func NewInstanceHttpServer(
	handlerFn http.Handler,
	opts ...httpServerOptions,
) IHTTPInstanceServer {
	hs := &httpServer{}
	hs.HttpContext = context.Background() // Set as default context
	hs.Addr = "0.0.0.0:8080"              // Set default port as 8080
	hs.Logger = logrus.StandardLogger()   // Set Logrus standard logger as default log

	// Set HTTP Handler
	srv := &http.Server{
		Handler: handlerFn,
	}
	hs.iHttpServer = srv

	for _, opt := range opts {
		opt(hs)
	}

	if hs.iHttpServer == nil {
		hs.Logger.Fatalf("http_server: Fatal! HTTP handler cannot be blank")
		os.Exit(1)
	}

	return hs
}

func (hs *httpServer) gracefullShutdown(quit chan os.Signal) {
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit
	hs.Logger.Warnf("Got signal: %v, shutting down server...", quit)

	ctx, cancel := context.WithTimeout(hs.HttpContext, 5*time.Second)
	defer cancel()

	hs.SetKeepAlivesEnabled(false)
	if err := hs.Shutdown(ctx); err != nil {
		hs.Logger.Errorf("Error when try to shutdown server. Err: %v", err)
	}

	hs.Logger.Infof("Shutting down server! Good bye")
}

func (hs *httpServer) Start(quit chan os.Signal) {
	// 	Run HTTP service
	go func() {
		if err := hs.ListenAndServe(); err != nil {
			hs.Logger.Fatalf("Fatal listening service! Err: %v", err)
		}
	}()

	hs.Logger.Infof("API service running on: %s", hs.Addr)

	hs.gracefullShutdown(quit)
}
