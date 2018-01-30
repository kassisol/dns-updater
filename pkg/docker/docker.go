package docker

import (
	"github.com/docker/docker/client"
)

var cnameLabel string = "dns.cname"
var hostLabel string = "dns.host"

func New() (*Config, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Config{
		Client: cli,
	}, nil
}
