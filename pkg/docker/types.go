package docker

import (
	"github.com/docker/docker/client"
)

type Config struct {
	Client *client.Client
}

type RecordCNAME struct {
	Name      string `json:"name"`
	Canonical string `json:"canonical"`
}
