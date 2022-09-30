package internal

import (
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertMessage(msg *md.Message, envoy *servicev3.StreamAccessLogsMessage) {
	// Convert node
	logs := envoy.GetHttpLogs()
	if logs != nil {
		for _, entry := range logs.GetLogEntry() {
			ce := new(md.CombinedEntry)
			if entry != nil {
			}
			msg.LogEntries = append(msg.LogEntries, ce)
		}
	}
	return
}
