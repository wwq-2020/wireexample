package home

import (
	"io/ioutil"
	"testing"

	"net/http/httptest"

	"github.com/wwq1988/wireexample/pkg/repo"
)

func TestHome(t *testing.T) {
	repo := &repo.Repo{}
	home := New(repo)

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	home.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != testdata {
		t.Errorf("expected:%s,got:%s", testdata, body)
	}
}
