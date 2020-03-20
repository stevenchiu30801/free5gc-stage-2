/*
 * NSSF Plugin
 */

package plugin

import (
	. "gofree5gc/lib/openapi/models"
)

type NetworkSliceManagementItem struct {
	SnssaiList []Snssai `json:"snssaiList"`

	PlmnIdList []PlmnId `json:"plmnIdList"`

	TaiList []Tai `json:"taiList"`

	NsiInformationList []NsiInformation `json:"nsiInformationList"`
}
