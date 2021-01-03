package service

import (
	"github.com/google/wire"
	"github.com/wljgithub/mall-project/internal/repository"
)

type Service struct {
	Repo repository.Repository
}






var Provider = wire.NewSet(repository.Provider, NewService)

func NewService(repo repository.Repository) (*Service, func(), error) {
	s := &Service{repo}
	return s, func() {}, nil
}
