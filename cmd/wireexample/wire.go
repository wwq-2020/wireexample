//+build wireinject

package main

import (
	"github.com/wwq1988/wireexample/pkg/conf"
	"github.com/wwq1988/wireexample/pkg/repo"
	"github.com/wwq1988/wireexample/pkg/server"

	"github.com/google/wire"
)

func Create() (*server.Server, error) {
	wire.Build(conf.New, repo.New, server.Set)
	return &server.Server{}, nil
}
