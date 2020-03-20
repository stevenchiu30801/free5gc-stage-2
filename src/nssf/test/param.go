/*
 * NSSF Testing Utility
 */

package test

import (
	"flag"

	"gofree5gc/lib/path_util"
	. "gofree5gc/src/nssf/plugin"
)

var (
	ConfigFileFromArgs string
	DefaultConfigFile  string = path_util.Gofree5gcPath("gofree5gc/src/nssf/test/conf/test_nssf_config.yaml")
)

type TestingUtil struct {
	ConfigFile string
}

type TestingNsselection struct {
	ConfigFile string

	GenerateNonRoamingQueryParameter func() NsselectionQueryParameter

	GenerateRoamingQueryParameter func() NsselectionQueryParameter
}

type TestingNssaiavailability struct {
	ConfigFile string

	NfId string

	SubscriptionId string

	NfNssaiAvailabilityUri string
}

func init() {
	flag.StringVar(&ConfigFileFromArgs, "config-file", DefaultConfigFile, "Configuration file")
	flag.Parse()
}
