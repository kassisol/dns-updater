package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func (c *Config) GetServicesCNAMEs() ([]RecordCNAME, error) {
	result := []RecordCNAME{}

	filterSL := filters.NewArgs()
	filterSL.Add("label", cnameLabel)
	filterSL.Add("label", hostLabel)

	listOpts := types.ServiceListOptions{
		Filters: filterSL,
	}

	services, err := c.Client.ServiceList(context.Background(), listOpts)
	if err != nil {
		return result, err
	}

	for _, service := range services {
		cname := RecordCNAME{
			Name:      service.Spec.Labels[cnameLabel],
			Canonical: service.Spec.Labels[hostLabel],
		}

		result = append(result, cname)
	}

	return result, nil
}

func (c *Config) GetServiceCNAME(name string) (RecordCNAME, error) {
	opts := types.ServiceInspectOptions{}

	service, _, err := c.Client.ServiceInspectWithRaw(context.Background(), name, opts)
	if err != nil {
		return RecordCNAME{}, err
	}

	if _, ok := service.Spec.Labels[cnameLabel]; !ok {
		return RecordCNAME{}, fmt.Errorf("%s label is not set", cnameLabel)
	}
	if _, ok := service.Spec.Labels[hostLabel]; !ok {
		return RecordCNAME{}, fmt.Errorf("%s label is not set", hostLabel)
	}

	cname := RecordCNAME{
		Name:      service.Spec.Labels[cnameLabel],
		Canonical: service.Spec.Labels[hostLabel],
	}

	return cname, nil
}
