package internal

import (
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertMessage(msg *md.Message, envoy *servicev3.StreamAccessLogsMessage) {
	if envoy == nil || msg == nil {
		return
	}
	msg.Dictionary = util.CreateInvertedDictionary(false)
	logs := envoy.GetHttpLogs()
	if logs != nil {
		for _, entry := range logs.GetLogEntry() {
			combined := new(md.CombinedEntry)
			combined.Http = ConvertHttpAccessLogEntry(msg.Dictionary, entry)
			combined.Common = ConvertCommon(msg.Dictionary, entry.GetCommonProperties())
			// TODO : convert traffic
			msg.LogEntries = append(msg.LogEntries, combined)
		}
	}
	logs2 := envoy.GetTcpLogs()
	if logs2 != nil {
		for _, entry := range logs2.GetLogEntry() {
			combined := new(md.CombinedEntry)
			combined.Tcp = ConvertTcpAccessLogEntry(entry)
			combined.Common = ConvertCommon(msg.Dictionary, entry.GetCommonProperties())
			// TODO : convert traffic
			msg.LogEntries = append(msg.LogEntries, combined)
		}
	}
}
