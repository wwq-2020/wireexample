package main

import (
	"flag"

	"github.com/wwq1988/wireexample/pkg/conf"
	"github.com/wwq1988/wireexample/pkg/grace"
)

func init() {
	flag.StringVar(&conf.Path, "conf", "conf.toml", "-conf=./conf")
	flag.Parse()
}

func main() {
	server, err := Create()
	if err != nil {
		panic(err)
	}
	grace.OnExit(server.Stop)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
