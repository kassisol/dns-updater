package main

import (
	"io"
	"os"

	"github.com/juliengk/go-utils/validation"
	"github.com/kassisol/dns-updater/pkg/dns"
	"github.com/kassisol/dns-updater/pkg/docker"
	"github.com/kassisol/dns-updater/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runApp(cmd *cobra.Command, args []string) {
	if cmdVersion {
		v := version.New()
		v.ShowVersion()
		os.Exit(0)
	}

	if cmdDebug {
		log.SetLevel(log.DebugLevel)
	}

	docker, err := docker.New()
	if err != nil {
		log.Fatal(err)
	}

	docker.Ping()

	d, err := dns.NewDriver(cmdDNSDriver, cmdConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	// Add CNAMEs for existing services at start
	log.Info("Checking for CNAMEs' of running services")

	cnames, err := docker.GetServicesCNAMEs()
	if err != nil {
		log.Fatal(err)
	}

	for _, cname := range cnames {
		log.WithFields(log.Fields{
			"name":      cname.Name,
			"canonical": cname.Canonical,
		}).Debug("Adding cname")

		if err = validation.IsValidFQDN(cname.Name); err != nil {
			log.Error(err)
			continue
		}
		if err = validation.IsValidFQDN(cname.Canonical); err != nil {
			log.Error(err)
			continue
		}

		if err := d.Add(cname.Name, cname.Canonical); err != nil {
			log.Error(err)
			continue
		}

		log.WithFields(log.Fields{
			"name":      cname.Name,
			"canonical": cname.Canonical,
		}).Info("Added cname")
	}

	// Events
	log.Info("Listening to Docker's events")

	messages, errs := docker.Events()

	for {
		select {
		case err := <-errs:
			if err != nil && err != io.EOF {
				log.Error(err)
			}

			os.Exit(1)
		case e := <-messages:
			cname, err := docker.GetServiceCNAME(e.Actor.ID)
			if err != nil {
				log.Error(err)
				continue
			}

			log.WithFields(log.Fields{
				"name":      cname.Name,
				"canonical": cname.Canonical,
			}).Debug("Adding cname")

			if err = validation.IsValidFQDN(cname.Name); err != nil {
				log.Error(err)
				continue
			}
			if err = validation.IsValidFQDN(cname.Canonical); err != nil {
				log.Error(err)
				continue
			}

			if err := d.Add(cname.Name, cname.Canonical); err != nil {
				log.Error(err)
				continue
			}

			log.WithFields(log.Fields{
				"name":      cname.Name,
				"canonical": cname.Canonical,
			}).Info("Added cname")
		}
	}
}
