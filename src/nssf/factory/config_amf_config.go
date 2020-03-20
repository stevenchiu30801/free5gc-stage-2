/*
 * NSSF Configuration Factory
 */

package factory

import (
	. "gofree5gc/lib/openapi/models"
)

type AmfConfig struct {
	NfId string `yaml:"nfId"`

	SupportedNssaiAvailabilityData []SupportedNssaiAvailabilityData `yaml:"supportedNssaiAvailabilityData"`
}
