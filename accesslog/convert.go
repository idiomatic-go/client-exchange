package accesslog

import (
	accesslogv32 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	accesslogv3 "github.com/idiomatic-go/client-exchange/accesslog/envoy"
	mdata "github.com/idiomatic-go/metric-data/accesslog"
)

func convert(msg *mdata.StreamAccessLogsMessage, envoy *accesslogv3.StreamAccessLogsMessage) {
	setTraffic(msg, envoy)
	//logs := msg.GetHttpLogs().
	// Create default view
	//convertDefaultView(log, msg)
	//convertAppIngress(log, msg)
	//convertAppEgress(log, msg)
	//convertAppIngress(log, msg)
	//convertAppEgress(log, msg)

}

func setTraffic(log *mdata.StreamAccessLogsMessage, envoy *accesslogv3.StreamAccessLogsMessage) {

}

func convertCommon(common *mdata.AccessLogCommon, envoy *accesslogv32.AccessLogCommon) {
	//common.FilterStateObjects = envoy.FilterStateObjects
	common.SampleRate = envoy.GetSampleRate()
	common.RouteName = envoy.GetRouteName()
	/*
		common.StartTime  = envoy.GetStartTime()
		common.Duration  = envoy..GetDu.Duration

		envoy.

		//common.DownstreamRemoteAddress  = envoy.DownstreamRemoteAddress
		//common.DownstreamLocalAddress = envoy.DownstreamLocalAddress *Address
		//common.TlsProperties  = envoy.TlsProperties*TLSProperties

		common.TimeToLastRxByte = envoy.GetTimeToLastRxByte()
		common.TimeToFirstUpstreamTxByte  = envoy.GetTimeToFirstUpstreamTxByte()
		common.TimeToLastUpstreamTxByte  = envoy.TimeToLastUpstreamTxByte*time.Duration
		common.TimeToFirstUpstreamRxByte  = envoy.TimeToFirstUpstreamRxByte*time.Duration
		common.TimeToLastUpstreamRxByte  = envoy.TimeToLastUpstreamRxByte*time.Duration
		common.TimeToFirstDownstreamTxByte  = envoy.TimeToFirstDownstreamTxByte*time.Duration
		common.TimeToLastDownstreamTxByte  = envoy.*time.Duration
		common.UpstreamRemoteAddress  = envoy.UpstreamRemoteAddress*Address
		common.UpstreamLocalAddress  = envoy.UpstreamLocalAddress*Address
		common.UpstreamCluster  = envoy.UpstreamClusterstring
		common.ResponseFlags  = envoy.ResponseFlags*ResponseFlags

		common.Metadata  = envoy.Metadata*v32.Metadata


		common.UpstreamTransportFailureReason = envoy.UpstreamTransportFailureReason
		common.RouteName  = envoy.RouteNamestring
		common.DownstreamDirectRemoteAddress  = envoy.DownstreamDirectRemoteAddress*Address


		common.CustomTags  = envoy.CustomTags  map[string]string

		common.UpstreamRequestAttemptCount  = envoy.UpstreamRequestAttemptCount

		common.ConnectionTerminationDetails  = envoy.ConnectionTerminationDetails string


	*/
}
