// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/wljgithub/mall-project/internal/api"
	"github.com/wljgithub/mall-project/internal/repository"
	"github.com/wljgithub/mall-project/internal/service"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	client, cleanup, err := repository.NewRedis()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := repository.NewMysql()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	repositoryRepository, cleanup3, err := repository.New(client, db)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceService, cleanup4, err := service.NewService(repositoryRepository)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, cleanup5, err := api.NewApiServer(serviceService)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app, cleanup6, err := NewApp(server)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
