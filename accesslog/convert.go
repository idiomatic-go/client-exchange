package accesslog

import (
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	accesslogv32 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	accesslogv3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
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
	//envoy.GetHttpLogs().GetLogEntry()[0].GetCommonProperties().
}

func convertCommon(log *mdata.AccessLogCommon, envoy *accesslogv32.AccessLogCommon) {
	log.SampleRate = envoy.GetSampleRate()
	log.RouteName = envoy.GetRouteName()
	log.UpstreamCluster = envoy.GetUpstreamCluster()
	//log.ConnectionTerminationDetails = envoy.GetConnectionTerminationDetails()
	log.UpstreamTransportFailureReason = envoy.GetUpstreamTransportFailureReason()
	if envoy.GetStartTime() != nil {
		t := envoy.GetStartTime().AsTime()
		log.StartTime = &t
	}
	//if envoy.GetDuration() != nil {
	//	t := envoy.GetDuration().AsTime()
	//	common.Duration = &t
	//}
	convertTls(log, envoy)
	convertResponseFlags(log, envoy)
	if envoy.GetTimeToLastRxByte() != nil {
		t := envoy.GetTimeToLastRxByte().AsDuration()
		log.TimeToLastRxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamTxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamTxByte().AsDuration()
		log.TimeToFirstUpstreamTxByte = &t
	}
	if envoy.GetTimeToLastUpstreamTxByte() != nil {
		t := envoy.GetTimeToLastUpstreamTxByte().AsDuration()
		log.TimeToLastUpstreamTxByte = &t
	}
	if envoy.GetTimeToFirstUpstreamRxByte() != nil {
		t := envoy.GetTimeToFirstUpstreamRxByte().AsDuration()
		log.TimeToFirstUpstreamRxByte = &t
	}
	if envoy.GetTimeToLastUpstreamRxByte() != nil {
		t := envoy.GetTimeToLastUpstreamRxByte().AsDuration()
		log.TimeToLastUpstreamRxByte = &t
	}
	if envoy.GetTimeToFirstDownstreamTxByte() != nil {
		t := envoy.GetTimeToFirstDownstreamTxByte().AsDuration()
		log.TimeToFirstDownstreamTxByte = &t
	}
	if envoy.GetTimeToLastDownstreamTxByte() != nil {
		t := envoy.GetTimeToLastDownstreamTxByte().AsDuration()
		log.TimeToLastDownstreamTxByte = &t
	}
	convertAddress(envoy.GetDownstreamRemoteAddress())
	/*

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

func convertAddress(envoy *corev3.Address) *mdata.Address {
	if envoy == nil {
		return nil
	}
	address := new(mdata.Address)
	if envoy.GetSocketAddress() != nil {

	} else {
		if envoy.GetPipe() != nil {

		} else {
			if envoy.GetEnvoyInternalAddress() != nil {

			}
		}
	}
	return address
}

func convertTls(log *mdata.AccessLogCommon, envoy *accesslogv32.AccessLogCommon) {
	if log == nil || envoy == nil || envoy.GetTlsProperties() == nil {
		return
	}
	log.TlsProperties = new(mdata.TLSProperties)
	//vers := strconv.Itoa(int(envoy.TlsProperties.TlsVersion))
	//log.TlsProperties.TlsVersion = log.Conmdata.TLSProperties_TLSVersion(strconv.Atoi(vers))

}

func convertResponseFlags(log *mdata.AccessLogCommon, envoy *accesslogv32.AccessLogCommon) {
	if log == nil || envoy == nil || envoy.GetResponseFlags() == nil {
		return
	}
	//log.

}
