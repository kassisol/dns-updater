package infoblox

import (
	"encoding/json"
	"fmt"

	"github.com/juliengk/stack/client"
)

func (c *Config) List(canonical string) (string, []string, error) {
	cc := &client.Config{
		Scheme: "https",
		Host:   c.IPAMHost,
		Port:   443,
		Path:   fmt.Sprintf("/wapi/v%s/record:host", c.APIVersion),
	}

	req, err := client.New(cc)
	if err != nil {
		return "", []string{}, err
	}

	req.HeaderAdd("Accept", "application/json")
	req.SetBasicAuth(c.Username, c.Password)
	req.ValueAdd("_return_fields", "aliases")
	req.ValueAdd("name", canonical)

	result := req.Get()

	if result.Response.StatusCode != 200 {
		return "", []string{}, fmt.Errorf("Something went wrong")
	}

	var response []RecordHost
	if err := json.Unmarshal(result.Body, &response); err != nil {
		return "", []string{}, err
	}

	if len(response) == 0 {
		return "", []string{}, nil
	}

	host := response[0]

	return host.Ref, host.Aliases, nil
}
