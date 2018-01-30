// DNS Updater is an application to add cname to DNS provider.
// Copyright (C) 2017 Kassisol inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cmdConfigFile string
	cmdDebug      bool
	cmdDNSDriver  string
	cmdVersion    bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "dns-updater",
		Short: "Manage DNS CNAME",
		Long:  "DNS Updater runs as a Docker service on a Docker Swarm manager host and listen for service create events then add cname to DNS provider based on label provided.",
		Run:   runApp,
	}

	rootCmd.Flags().StringVarP(&cmdConfigFile, "config", "c", "config.yml", "Config file")
	rootCmd.Flags().BoolVarP(&cmdDebug, "debug", "D", false, "Enable debug mode")
	rootCmd.Flags().StringVarP(&cmdDNSDriver, "driver", "d", "infoblox", "DNS Driver")
	rootCmd.Flags().BoolVarP(&cmdVersion, "version", "v", false, "Show the DNS Updater version information")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
