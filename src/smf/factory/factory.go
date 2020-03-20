/*
 * AMF Configuration Factory
 */

package factory

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"gofree5gc/src/smf/logger"
)

var SmfConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.AppLog.Fatal(err)
	}
}

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	SmfConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &SmfConfig)
	checkErr(err)

	logger.InitLog.Infof("Successfully initialize configuration %s", f)
}