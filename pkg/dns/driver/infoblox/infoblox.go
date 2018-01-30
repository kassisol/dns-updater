package infoblox

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/dns-updater/pkg/dns"
	"github.com/kassisol/dns-updater/pkg/dns/driver"
	"gopkg.in/yaml.v2"
)

func init() {
	dns.RegisterDriver("infoblox", New)
}

func New(config string) (driver.Storager, error) {
	apiVersionRe := regexp.MustCompile(`[0-9]+\.[0-9]+`)

	data, err := ioutil.ReadFile(config)
	if err != nil {
		return nil, err
	}

	c := new(Config)

	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	if len(c.IPAMHost) == 0 {
		return nil, fmt.Errorf("Empty or missing key 'ipam_host' in config file")
	}

	if err = validation.IsValidFQDN(c.IPAMHost); err != nil {
		return nil, err
	}

	if len(c.APIVersion) == 0 {
		return nil, fmt.Errorf("Empty or missing key 'api_versiont' in config file")
	}

	if !apiVersionRe.MatchString(c.APIVersion) {
		return nil, fmt.Errorf("API version is not valid")
	}

	if len(c.Username) == 0 {
		return nil, fmt.Errorf("Empty or missing key 'username' in config file")
	}

	if len(c.Password) == 0 {
		return nil, fmt.Errorf("Empty or missing key 'password' in config file")
	}

	return c, nil
}
