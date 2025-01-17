/*
 * NSSF Configuration Factory
 */

package factory

import (
	. "gofree5gc/lib/openapi/models"
)

type MappingFromPlmnConfig struct {
	OperatorName string `yaml:"operatorName,omitempty"`

	HomePlmnId *PlmnId `yaml:"homePlmnId"`

	MappingOfSnssai []MappingOfSnssai `yaml:"mappingOfSnssai"`
}
