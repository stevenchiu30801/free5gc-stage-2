//+build debug

package ausf_util

import (
	"gofree5gc/lib/path_util"
)

var AusfLogPath = path_util.Gofree5gcPath("gofree5gc/ausfsslkey.log")
var AusfPemPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/ausf.pem")
var AusfKeyPath = path_util.Gofree5gcPath("gofree5gc/support/TLS/ausf.key")
var DefaultAusfConfigPath = path_util.Gofree5gcPath("gofree5gc/config/ausfcfg.conf")
