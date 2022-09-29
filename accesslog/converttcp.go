package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/metric-data/accesslog"
)

func convertTcpEntry(envoy *datav3.TCPAccessLogEntry) *accesslog.TCPAccessLogEntry {
	if envoy == nil || envoy.GetConnectionProperties() == nil {
		return nil
	}
	tcp := new(accesslog.TCPAccessLogEntry)
	tcp.ConnectionProperties.SentBytes = envoy.GetConnectionProperties().GetSentBytes()
	tcp.ConnectionProperties.ReceivedBytes = envoy.GetConnectionProperties().GetReceivedBytes()
	return tcp
}
