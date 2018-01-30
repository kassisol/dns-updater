package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
)

func (c *Config) Events() (<-chan events.Message, <-chan error) {
	filterE := filters.NewArgs()
	filterE.Add("type", events.ServiceEventType)

	eventsOpts := types.EventsOptions{
		Filters: filterE,
	}

	return c.Client.Events(context.Background(), eventsOpts)
}
