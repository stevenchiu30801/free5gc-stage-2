package pcf_util

import (
	"encoding/hex"
	"gofree5gc/lib/openapi/models"
	"gofree5gc/src/pcf/factory"
	"gofree5gc/src/pcf/logger"
	"gofree5gc/src/pcf/pcf_context"

	"github.com/google/uuid"
)

// Init PCF Context from config flie
func InitpcfContext(context *pcf_context.PCFContext) {
	config := factory.PcfConfig
	logger.UtilLog.Infof("pcfconfig Info: Version[%s] Description[%s]", config.Info.Version, config.Info.Description)
	configuration := config.Configuration
	context.NfId = uuid.New().String()
	if configuration.PcfName != "" {
		context.Name = configuration.PcfName
	}
	sbi := configuration.Sbi
	context.NrfUri = configuration.NrfUri
	context.UriScheme = models.UriScheme(sbi.Scheme)
	context.HttpIPv4Address = "127.0.0.1" // default localhost
	context.HttpIpv4Port = 29507          // default port
	if sbi != nil {
		if sbi.IPv4Addr != "" {
			context.HttpIPv4Address = sbi.IPv4Addr
		}
		if sbi.Port != 0 {
			context.HttpIpv4Port = sbi.Port
		}
		if sbi.Scheme == "https" {
			context.UriScheme = models.UriScheme_HTTPS
		} else {
			context.UriScheme = models.UriScheme_HTTP
		}
	}
	serviceList := configuration.ServiceList
	context.InitNFService(serviceList, config.Info.Version)
	context.TimeFormat = configuration.TimeFormat
	context.DefaultBdtRefId = configuration.DefaultBdtRefId
	for _, service := range context.NfService {
		context.PcfServiceUris[service.ServiceName] = service.ApiPrefix + "/" + string(service.ServiceName) + "/" + (*service.Versions)[0].ApiVersionInUri
		context.PcfSuppFeats[service.ServiceName], _ = hex.DecodeString(service.SupportedFeatures)
	}
}
