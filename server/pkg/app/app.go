package app

import (
	"context"
	"github.com/wljgithub/mall-project/internal/api"
	"github.com/wljgithub/mall-project/pkg/conf"
	"github.com/wljgithub/mall-project/pkg/log"
	"golang.org/x/sync/errgroup"
	"time"
)

type App struct {
	eg     *errgroup.Group
	Server api.Server
}

func NewApp(server api.Server) (app *App, cancelFunc func(), err error) {
	app = &App{Server: server, eg: new(errgroup.Group)}
	cancelFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//
		//for _, server := range servers {
		if err := server.Shutdown(ctx); err != nil {
			log.Info(err)
		}
		//}
	}
	return
}

func (this *App) Run() {
	//for _, server := range this.Server {
	this.eg.Go(func() error {
		log.Info("server running on", conf.Conf.App.Addr)
		return this.Server.Run()
	})
	//}

}
func (this *App) Wait() error {
	return this.eg.Wait()
}
