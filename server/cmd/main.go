package main

import (
	"flag"
	"github.com/wljgithub/mall-project/pkg/app"
	"github.com/wljgithub/mall-project/pkg/conf"
	"github.com/wljgithub/mall-project/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

var path = flag.String("config", "conf/config.local.yml", "config path")

func main() {
	flag.Parse()

	// init config
	if err := conf.Init(*path); err != nil {
		panic(err)
	}
	//init log
	if err := log.Init(); err != nil {
		panic(err)
	}
	// init app
	app, cancel, err := app.InitApp()
	if err != nil {
		//log.Fatalf("%+v", err)
		panic(err)
	}
	// app run
	app.Run()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	switch <-quit {
	case syscall.SIGTERM, syscall.SIGINT:
		// app cancel
		cancel()
		log.Info("receive kill,server exit")
	case syscall.SIGHUP:
		log.Info("receive hangup signal")
	}
	// app wait
	app.Wait()
}
