/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gofree5gc/src/app"
	"gofree5gc/src/smf/logger"
	"gofree5gc/src/smf/smf_service"
	"os"
)

var SMF = &smf_service.SMF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "smf"
	fmt.Print(app.Name, "\n")
	app.Usage = "-free5gccfg common configuration file -smfcfg smf configuration file"
	app.Action = action
	app.Flags = SMF.GetCliCmd()

	if err := app.Run(os.Args); err != nil {
		logger.AppLog.Errorf("SMF Run error: %v", err)
	}
}

func action(c *cli.Context) {
	app.AppInitializeWillInitialize(c.String("free5gccfg"))
	SMF.Initialize(c)
	SMF.Start()
}
