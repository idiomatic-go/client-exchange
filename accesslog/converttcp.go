package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func convertTcpEntry(envoy *datav3.TCPAccessLogEntry) *md.TCPAccessLogEntry {
	if envoy == nil || envoy.GetConnectionProperties() == nil {
		return nil
	}
	tcp := new(md.TCPAccessLogEntry)
	tcp.ConnectionProperties = new(md.ConnectionProperties)
	tcp.ConnectionProperties.SentBytes = envoy.GetConnectionProperties().GetSentBytes()
	tcp.ConnectionProperties.ReceivedBytes = envoy.GetConnectionProperties().GetReceivedBytes()
	return tcp
}
