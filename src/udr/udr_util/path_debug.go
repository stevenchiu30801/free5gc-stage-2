//+build debug

package udr_util

import (
	"gofree5gc/lib/path_util"
)

var UdrLogPath = path_util.Gofree5gcPath("gofree5gc/udrsslkey.log")
var UdrPemPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/_debug.pem")
var UdrKeyPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/_debug.key")
var DefaultUdrConfigPath = path_util.Gofree5gcPath("gofree5gc/config/udrcfg.conf")
