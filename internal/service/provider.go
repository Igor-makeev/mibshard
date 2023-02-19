package service

import (
	"net/http"
	"time"
)

type Provider struct {
	Client *http.Client
}

func NewProvider() *Provider {
	client := &http.Client{}
	transport := &http.Transport{}
	transport.MaxIdleConns = 20
	client.Transport = transport
	client.Timeout = time.Second * 1
	return &Provider{
		Client: client,
	}
}
