package internal

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertHttpAccessLogEntry(dictionary *util.InvertedDictionary, envoy *datav3.HTTPAccessLogEntry) *md.HTTPAccessLogEntry {
	if envoy == nil {
		return nil
	}
	entry := new(md.HTTPAccessLogEntry)
	entry.ProtocolVersion = md.HTTPAccessLogEntry_HTTPVersion(envoy.GetProtocolVersion())
	if envoy.GetRequest() != nil {
		entry.Request = ConvertHttpRequest(dictionary, envoy.GetRequest())
	}
	if envoy.GetResponse() != nil {
		entry.Response = ConvertHttpResponse(dictionary, envoy.GetResponse())
	}
	return entry
}

func ConvertHttpRequest(dictionary *util.InvertedDictionary, envoy *datav3.HTTPRequestProperties) *md.HTTPRequestProperties {
	if envoy == nil {
		return nil
	}
	req := new(md.HTTPRequestProperties)
	req.RequestMethod = md.RequestMethod(envoy.GetRequestMethod())
	req.Scheme = dictionary.Add(envoy.GetScheme())
	req.Authority = dictionary.Add(envoy.GetAuthority())
	if envoy.GetPort() != nil {
		req.Port = envoy.GetPort().Value
	}
	req.Path = dictionary.Add(envoy.GetPath())
	req.UserAgent = dictionary.Add(envoy.GetUserAgent())
	req.Referer = dictionary.Add(envoy.GetReferer())
	req.ForwardedFor = dictionary.Add(envoy.GetForwardedFor())
	req.RequestId = envoy.GetRequestId()
	req.OriginalPath = dictionary.Add(envoy.GetOriginalPath())
	req.RequestHeadersBytes = envoy.GetRequestHeadersBytes()
	req.RequestBodyBytes = envoy.GetRequestBodyBytes()
	req.RequestHeaders = envoy.GetRequestHeaders()
	return req
}

func ConvertHttpResponse(dictionary *util.InvertedDictionary, envoy *datav3.HTTPResponseProperties) *md.HTTPResponseProperties {
	if envoy == nil {
		return nil
	}
	resp := new(md.HTTPResponseProperties)
	if envoy.GetResponseCode() != nil {
		resp.ResponseCode = envoy.GetResponseCode().Value
	}
	resp.ResponseHeadersBytes = envoy.GetResponseHeadersBytes()
	resp.ResponseBodyBytes = envoy.GetResponseBodyBytes()
	resp.ResponseHeaders = envoy.GetResponseHeaders()
	resp.ResponseTrailers = envoy.GetResponseTrailers()
	resp.ResponseCodeDetails = dictionary.Add(envoy.GetResponseCodeDetails())
	return resp
}
