package test

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var TestConfig RanConfig

type RanConfig struct {
	IpAddr string `yaml:"ipAddr"`

	AmfAddr string `yaml:"amfAddr"`

	MongoAddr string `yaml:"mongoAddr"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func ParseConfig(cfg string) {
	content, err := ioutil.ReadFile(cfg)
	checkErr(err)

	err = yaml.Unmarshal(content, &TestConfig)
	checkErr(err)
}
