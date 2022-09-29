package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
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

func convertHttpRequest(envoy *datav3.HTTPRequestProperties) *accesslog.HTTPRequestProperties {
	if envoy == nil {
		return nil
	}
	req := new(accesslog.HTTPRequestProperties)
	// mdata.SetRequestMethodName(int32(envoy.RequestMethod))
	//if envoy.GetRequestMethod()G
	req.Scheme = envoy.GetRequestId()
	req.Authority = envoy.GetAuthority()
	if envoy.GetPort() != nil {
		req.Port = envoy.GetPort().Value
	}
	req.Path = envoy.GetPath()
	req.UserAgent = envoy.GetUserAgent()
	req.Referer = envoy.GetReferer()
	req.ForwardedFor = envoy.GetForwardedFor()
	req.RequestId = envoy.GetRequestId()
	req.OriginalPath = envoy.GetOriginalPath()
	req.RequestHeadersBytes = envoy.GetRequestHeadersBytes()
	req.RequestBodyBytes = envoy.GetRequestBodyBytes()
	req.RequestHeaders = envoy.GetRequestHeaders()
	return req
}
