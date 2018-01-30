package infoblox

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/juliengk/go-utils"
	"github.com/juliengk/stack/client"
)

func (c *Config) Add(name, canonical string) error {
	ref, aliases, err := c.List(canonical)
	if err != nil {
		return err
	}

	if utils.StringInSlice(name, aliases, false) {
		return nil
	}

	aliases = append(aliases, name)

	cc := &client.Config{
		Scheme: "https",
		Host:   c.IPAMHost,
		Port:   443,
		Path:   fmt.Sprintf("/wapi/v%s/%s", c.APIVersion, ref),
	}

	req, err := client.New(cc)
	if err != nil {
		return err
	}

	req.HeaderAdd("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	newcname := UpdateRecordHostAliases{
		Aliases: aliases,
	}

	data, err := json.Marshal(newcname)
	if err != nil {
		return err
	}

	result := req.Put(bytes.NewBuffer(data))

	if result.Response.StatusCode != 200 {
		return fmt.Errorf("Something went wrong")
	}

	return nil
}
