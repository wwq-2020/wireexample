package mux

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wwq1988/wireexample/pkg/handlers/home"
	"github.com/wwq1988/wireexample/pkg/repo"
)

func New(repo *repo.Repo) http.Handler {
	r := mux.NewRouter().StrictSlash(true).SkipClean(true)
	sub := r.PathPrefix("/api").Subrouter()
	home.Inject(sub, repo)
	return r
}
