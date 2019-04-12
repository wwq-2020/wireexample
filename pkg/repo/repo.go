package repo

import (
	"github.com/wwq1988/wireexample/pkg/conf"
)

type Repo struct {
}

func New(conf *conf.Conf) (*Repo, error) {
	return &Repo{}, nil
}

func (rp *Repo) Close() {

}
