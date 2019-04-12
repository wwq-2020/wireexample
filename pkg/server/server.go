package server

import (
	"context"
	"net"
	"net/http"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/wwq1988/wireexample/pkg/conf"
	"github.com/wwq1988/wireexample/pkg/mux"
	"github.com/wwq1988/wireexample/pkg/repo"
)

var Set = wire.NewSet(New, mux.New)

type Server struct {
	listener net.Listener
	server   *http.Server
	repo     *repo.Repo
}

func New(conf *conf.Conf, handler http.Handler, repo *repo.Repo) (*Server, error) {
	listener, err := net.Listen(conf.ListenProto, conf.ListenAddr)
	if err != nil {
		return nil, err
	}
	server := &http.Server{Handler: handler}
	return &Server{
		listener: listener,
		server:   server,
		repo:     repo,
	}, nil
}

func (s *Server) Start() error {
	if err := s.server.Serve(s.listener); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *Server) Stop() {
	s.repo.Close()
	s.server.Shutdown(context.Background())
}
