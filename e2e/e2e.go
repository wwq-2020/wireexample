package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/wwq1988/wireexample/pkg/conf"
)

const (
	urlBase = "http://127.0.0.1:%s/api/%s"
)

func main() {
	conf.Path = "../cmd/wireexample/conf.toml"
	conf, err := conf.New()
	if err != nil {
		panic(err)
	}
	parts := strings.Split(conf.ListenAddr, ":")
	if len(parts) != 2 {
		panic("err addr fmt")
	}
	portStr := parts[1]
	url := fmt.Sprintf(urlBase, portStr, "home")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	expected := "helloworld"
	dataStr := string(data)
	if dataStr != expected {
		panic(fmt.Sprintf("expected:%s,got:%s", expected, dataStr))
	}

}
