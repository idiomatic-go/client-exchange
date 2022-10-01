package internal

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
	"time"
)

func ConvertCommon(traffic md.Common_Traffic, config *data.Configuration, dictionary *util.InvertedDictionary, envoy *datav3.AccessLogCommon) *md.CommonProperties {
	if envoy == nil {
		return nil
	}
	common := new(md.CommonProperties)
	common.Traffic = traffic
	common.SampleRate = envoy.GetSampleRate()
	common.DownstreamRemoteAddress = ConvertAddress(envoy.GetDownstreamRemoteAddress())
	common.DownstreamLocalAddress = ConvertAddress(envoy.GetDownstreamLocalAddress())

	common.TlsProperties = ConvertTls(envoy.GetTlsProperties())
	common.StartTime = ConvertTimestamp(envoy.GetStartTime())

	common.TimeToLastRxByte = ConvertDuration(envoy.GetTimeToLastRxByte())

	common.TimeToFirstUpstreamTxByte = ConvertDuration(envoy.GetTimeToFirstUpstreamTxByte())
	common.TimeToLastUpstreamTxByte = ConvertDuration(envoy.GetTimeToLastUpstreamTxByte())

	common.TimeToFirstUpstreamRxByte = ConvertDuration(envoy.GetTimeToFirstUpstreamRxByte())
	common.TimeToLastUpstreamRxByte = ConvertDuration(envoy.GetTimeToLastUpstreamRxByte())

	common.TimeToFirstDownstreamTxByte = ConvertDuration(envoy.GetTimeToFirstDownstreamTxByte())
	common.TimeToLastDownstreamTxByte = ConvertDuration(envoy.GetTimeToLastDownstreamTxByte())

	common.UpstreamRemoteAddress = ConvertAddress(envoy.GetUpstreamRemoteAddress())
	common.UpstreamLocalAddress = ConvertAddress(envoy.GetUpstreamLocalAddress())
	common.UpstreamCluster = dictionary.Add(envoy.GetUpstreamCluster())

	common.ResponseFlags = ConvertResponseFlags(envoy.GetResponseFlags())
	common.UpstreamTransportFailureReason = envoy.GetUpstreamTransportFailureReason()
	common.RouteName = dictionary.Add(envoy.GetRouteName())
	common.DownstreamDirectRemoteAddress = ConvertAddress(envoy.GetDownstreamRemoteAddress())

	if traffic == md.Common_Traffic_Egress {
		common.CustomTags = util.MapSubset(config.Egress.Custom, envoy.GetCustomTags())
	} else {
		common.CustomTags = util.MapSubset(config.Ingress.Custom, envoy.GetCustomTags())
	}

	// BUG : v3 common still the same as v2
	//if envoy.GetDuration() != nil {
	//	t := envoy.GetDuration().AsTime()
	//	common.Duration = &t
	//}
	//common.UpstreamRequestAttemptCount  = envoy.UpstreamRequestAttemptCount
	//common.ConnectionTerminationDetails = envoy.GetConnectionTerminationDetails()
	return common
}

func ConvertTimestamp(ts *timestamp.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

func ConvertDuration(d *duration.Duration) *time.Duration {
	if d == nil {
		return nil
	}
	t := d.AsDuration()
	return &t
}
