/*
 * NSSF Configuration Factory
 */

package factory

import (
	. "gofree5gc/lib/openapi/models"
)

type NsiConfig struct {
	Snssai *Snssai `yaml:"snssai"`

	NsiInformationList []NsiInformation `yaml:"nsiInformationList"`
}
