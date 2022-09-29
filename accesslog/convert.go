package accesslog

import (
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/idiomatic-go/metric-data/accesslog"
)

func convert(msg *accesslog.Message, envoy *servicev3.StreamAccessLogsMessage) {
	setTraffic(msg, envoy)
	//logs := msg.GetHttpLogs().
	// Create default view
	//convertDefaultView(log, msg)
	//convertAppIngress(log, msg)
	//convertAppEgress(log, msg)
	//convertAppIngress(log, msg)
	//convertAppEgress(log, msg)

}

func setTraffic(log *accesslog.Message, envoy *servicev3.StreamAccessLogsMessage) {
	//envoy.GetHttpLogs().GetLogEntry()[0].GetCommonProperties().
}
