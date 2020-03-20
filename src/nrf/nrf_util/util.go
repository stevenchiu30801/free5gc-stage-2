//+build !debug

package nrf_util

import (
	"gofree5gc/lib/path_util"
)

// Path of HTTP2 key and log file

var NrfLogPath = path_util.Gofree5gcPath("gofree5gc/nrfsslkey.log")
var NrfPemPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/nrf.pem")
var NrfKeyPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/nrf.key")
