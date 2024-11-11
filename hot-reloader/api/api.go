package api

import (
	"log"
	"net/http"

	"hot_reloader/config"
)

type Api struct {
	client *http.Client
	url    string
}

func New(url string) *Api {
	return &Api{
		client: http.DefaultClient,
		url:    url,
	}
}

func (a *Api) MakeRequest() int {
	res, err := a.client.Get(a.url)
	if err != nil {
		log.Println("failed to make request", err, "url", a.url)
		return 500
	}

	return res.StatusCode
}

func (a *Api) Reload(cfg any) {
	newURL := cfg.(*config.Config).URL
	a.url = newURL
}
