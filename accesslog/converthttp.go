package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func convertHttpRequest(envoy *datav3.HTTPRequestProperties) *md.HTTPRequestProperties {
	if envoy == nil {
		return nil
	}
	req := new(md.HTTPRequestProperties)
	req.RequestMethod = md.RequestMethod(envoy.GetRequestMethod())
	req.Scheme = envoy.GetScheme()
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

func convertHttpResponse(envoy *datav3.HTTPResponseProperties) *md.HTTPResponseProperties {
	if envoy == nil {
		return nil
	}
	resp := new(md.HTTPResponseProperties)
	resp.ResponseCode = envoy.GetResponseCode().Value
	resp.ResponseHeadersBytes = envoy.GetResponseHeadersBytes()
	resp.ResponseBodyBytes = envoy.GetResponseBodyBytes()
	resp.ResponseHeaders = envoy.GetResponseHeaders()
	resp.ResponseTrailers = envoy.GetResponseTrailers()
	resp.ResponseCodeDetails = envoy.GetResponseCodeDetails()
	return resp
}
