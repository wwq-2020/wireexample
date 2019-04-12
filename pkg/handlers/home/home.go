package home

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wwq1988/wireexample/pkg/repo"
)

const (
	path = "/home"
)

var (
	testdata = "helloword"
)

type Handler struct {
	repo *repo.Repo
}

func New(repo *repo.Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, testdata)
}

func (h *Handler) Path() string {
	return path
}

func Inject(r *mux.Router, repo *repo.Repo) {
	r.Handle(path, New(repo))
}
