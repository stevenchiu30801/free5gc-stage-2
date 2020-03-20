/*
 * NSSF
 *
 * Network Slice Selection Function
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"fmt"
	"gofree5gc/src/app"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"gofree5gc/src/nssf/logger"
	"gofree5gc/src/nssf/nssf_service"
)

var NSSF = &nssf_service.NSSF{}

var appLog *logrus.Entry

func init() {
	appLog = logger.AppLog
}

func main() {
	app := cli.NewApp()
	app.Name = "nssf"
	fmt.Print(app.Name, "\n")
	app.Usage = "-free5gccfg common configuration file -nssfcfg nssf configuration file"
	app.Action = action
	app.Flags = NSSF.GetCliCmd()

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("NSSF Run error: %v", err)
	}

}

func action(c *cli.Context) {
	app.AppInitializeWillInitialize(c.String("free5gccfg"))
	NSSF.Initialize(c)
	NSSF.Start()
}
