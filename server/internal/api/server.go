package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wljgithub/mall-project/internal/service"
	"github.com/wljgithub/mall-project/pkg/conf"
	"net/http"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type HttpServer struct {
	server *http.Server
	srv    *service.Service
}

var _ Server = &HttpServer{}

var Provider = wire.NewSet(service.Provider, NewApiServer)

func NewApiServer(s *service.Service) (Server, func(), error) {
	server, err := newApiServer(conf.Conf.App)
	if err != nil {
		return nil, func() {}, err
	}
	apiServer := &HttpServer{srv: s, server: server}
	apiServer.Init()
	return apiServer, func() {}, nil
}

func newApiServer(config conf.AppConfig) (*http.Server, error) {
	gin.SetMode(config.RunMode)
	return &http.Server{
		Addr: config.Addr,
	}, nil
}
func (this *HttpServer) Run() error {
	return this.server.ListenAndServe()
}

func (this *HttpServer) Shutdown(ctx context.Context) error {
	return this.server.Shutdown(ctx)
}

func (this *HttpServer) Init() {
	this.server.Handler = this.Load(gin.Default())
}





