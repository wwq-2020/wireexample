package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

var Path string

type Conf struct {
	ListenProto string
	ListenAddr  string
}

func New() (*Conf, error) {
	conf := &Conf{}
	if _, err := toml.DecodeFile(Path, conf); err != nil {
		return nil, errors.WithStack(err)
	}
	return conf, nil
}
