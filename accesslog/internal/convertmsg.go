package internal

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
	"strings"
)

func ConvertMessage(config *data.Configuration, msg *md.Message, envoy *servicev3.StreamAccessLogsMessage) {
	if envoy == nil || msg == nil {
		return
	}
	if msg.Dictionary == nil {
		msg.Dictionary = util.CreateInvertedDictionary(false)
	}
	if config == nil {
		config = new(data.Configuration)
	}
	logs := envoy.GetHttpLogs()
	if logs != nil {
		for _, entry := range logs.GetLogEntry() {
			traffic := CreateHttpTraffic(msg.Identifier.Application, entry)
			combined := new(md.CombinedEntry)
			combined.Http = ConvertHttpAccessLogEntry(traffic, config, msg.Dictionary, entry)
			combined.Common = ConvertCommon(traffic, config, msg.Dictionary, entry.GetCommonProperties())
			msg.LogEntries = append(msg.LogEntries, combined)
		}
	}
	logs2 := envoy.GetTcpLogs()
	if logs2 != nil {
		for _, entry := range logs2.GetLogEntry() {
			traffic := CreateTcpTraffic(entry)
			combined := new(md.CombinedEntry)
			combined.Tcp = ConvertTcpAccessLogEntry(entry)
			combined.Common = ConvertCommon(traffic, config, msg.Dictionary, entry.GetCommonProperties())
			msg.LogEntries = append(msg.LogEntries, combined)
		}
	}
}

func CreateHttpTraffic(app string, entry *datav3.HTTPAccessLogEntry) md.Common_Traffic {
	// if the authority is local host or matches the service name then it is ingress
	if strings.HasPrefix(entry.Request.Authority, "localhost") || strings.HasPrefix(entry.Request.Authority, app) {
		return md.Common_Traffic_Ingress
	}
	return md.Common_Traffic_Egress
}

func CreateTcpTraffic(entry *datav3.TCPAccessLogEntry) md.Common_Traffic {

	return md.Common_Traffic_UNSPECIFIED
}
