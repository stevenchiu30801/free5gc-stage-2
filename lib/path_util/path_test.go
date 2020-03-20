package path_util

import (
	"gofree5gc/lib/path_util/logger"
	"testing"
)

func TestFree5gcPath(t *testing.T) {
	logger.PathLog.Infoln(Gofree5gcPath("gofree5gc/abcdef/abcdef.pem"))
}
