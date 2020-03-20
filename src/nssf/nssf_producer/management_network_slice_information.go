/*
 * NSSF Management
 *
 * NSSF Management Service
 */

package nssf_producer

import (
	"net/http"
	"reflect"

	. "gofree5gc/lib/openapi/models"
	"gofree5gc/src/nssf/factory"
	. "gofree5gc/src/nssf/plugin"
	"gofree5gc/src/nssf/util"
)

// Management network slice information POST method
func networkSliceManagementPost(m NetworkSliceManagementDocument, d *ProblemDetails) (status int) {
	nssfConfig := factory.NssfConfig.Configuration

	var found bool

	for _, n := range m {
		// Create supported NSSAI in PLMN list
		for _, plmnId := range n.PlmnIdList {
			found = false
			for j, supportedNssaiInPlmn := range nssfConfig.SupportedNssaiInPlmnList {
				if reflect.DeepEqual(*supportedNssaiInPlmn.PlmnId, plmnId) {
					found = true
					// Add missing supported S-NSSAI in PLMN
					for _, snssai := range n.SnssaiList {
						if !util.Contain(snssai, supportedNssaiInPlmn.SupportedSnssaiList) {
							nssfConfig.SupportedNssaiInPlmnList[j].SupportedSnssaiList = append(nssfConfig.SupportedNssaiInPlmnList[j].SupportedSnssaiList, snssai)
						}
					}
					break
				}
			}

			if !found {
				// Create new SupportedNssaiInPlmn
				supportedNssaiInPlmn := factory.SupportedNssaiInPlmn{
					PlmnId:              &plmnId,
					SupportedSnssaiList: n.SnssaiList,
				}
				nssfConfig.SupportedNssaiInPlmnList = append(nssfConfig.SupportedNssaiInPlmnList, supportedNssaiInPlmn)
			}
		}

		// Create NSI
		for _, snssai := range n.SnssaiList {
			found = false
			for j, nsiConfig := range nssfConfig.NsiList {
				if reflect.DeepEqual(*nsiConfig.Snssai, snssai) {
					found = true
					// add missing NSI information
					for _, nsiInformation := range n.NsiInformationList {
						if !util.Contain(nsiInformation, nsiConfig.NsiInformationList) {
							nssfConfig.NsiList[j].NsiInformationList = append(nssfConfig.NsiList[j].NsiInformationList, nsiInformation)
						}
					}
					break
				}
			}

			if !found {
				// Create new NsiConfig
				nsiConfig := factory.NsiConfig{
					Snssai:             &snssai,
					NsiInformationList: n.NsiInformationList,
				}
				nssfConfig.NsiList = append(nssfConfig.NsiList, nsiConfig)
			}
		}

		// Create TA
		for _, tai := range n.TaiList {
			found = false
			for j, taConfig := range nssfConfig.TaList {
				if reflect.DeepEqual(*taConfig.Tai, tai) {
					found = true
					// add missing supported S-NSSAI in TAI
					for _, snssai := range n.SnssaiList {
						if !util.Contain(snssai, taConfig.SupportedSnssaiList) {
							nssfConfig.TaList[j].SupportedSnssaiList = append(nssfConfig.TaList[j].SupportedSnssaiList, snssai)
						}
					}
					break
				}
			}

			if !found {
				// Create new TaConfig
				var accessType AccessType = AccessType__3_GPP_ACCESS
				taConfig := factory.TaConfig{
					Tai:                 &tai,
					AccessType:          &accessType,
					SupportedSnssaiList: n.SnssaiList,
				}
				nssfConfig.TaList = append(nssfConfig.TaList, taConfig)
			}
		}
	}

	status = http.StatusCreated
	return
}
