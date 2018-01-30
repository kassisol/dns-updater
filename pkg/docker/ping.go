package docker

import (
	"context"
	"fmt"
	"time"
)

func (c *Config) Ping() {
	for {
		if _, err := c.Client.Ping(context.Background()); err == nil {
			break
		} else {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 5)
	}
}
