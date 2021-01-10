package repository

import "github.com/google/wire"

func InitRepository() (*Repo, func(), error) {
	panic(wire.Build(Provider))
}
