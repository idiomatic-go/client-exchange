package internal

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertHttpAccessLogEntry(traffic md.Common_Traffic, config *data.Configuration, dictionary *util.InvertedDictionary, envoy *datav3.HTTPAccessLogEntry) *md.HTTPAccessLogEntry {
	if envoy == nil {
		return nil
	}
	entry := new(md.HTTPAccessLogEntry)
	entry.ProtocolVersion = md.HTTPAccessLogEntry_HTTPVersion(envoy.GetProtocolVersion())
	if envoy.GetRequest() != nil {
		entry.Request = ConvertHttpRequest(traffic, config, dictionary, envoy.GetRequest())
	}
	if envoy.GetResponse() != nil {
		entry.Response = ConvertHttpResponse(traffic, config, dictionary, envoy.GetResponse())
	}
	return entry
}

func ConvertHttpRequest(traffic md.Common_Traffic, config *data.Configuration, dictionary *util.InvertedDictionary, envoy *datav3.HTTPRequestProperties) *md.HTTPRequestProperties {
	if envoy == nil {
		return nil
	}
	if config == nil {
		config = new(data.Configuration)
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
	if traffic == md.Common_Traffic_Egress {
		req.RequestHeaders = util.MapSubset(config.Egress.RequestHeaders, envoy.GetRequestHeaders())
		// TODO - update
		//req.Cookies = ConvertCookies(config.Egress.Cookies,envoy.GetRequestHeaders())
	} else {
		req.RequestHeaders = util.MapSubset(config.Ingress.RequestHeaders, envoy.GetRequestHeaders())
		// TODO - update
		//req.Cookies = ConvertCookies(config.Ingress.Cookies,envoy.GetRequestHeaders())
	}
	return req
}

func ConvertHttpResponse(traffic md.Common_Traffic, config *data.Configuration, dictionary *util.InvertedDictionary, envoy *datav3.HTTPResponseProperties) *md.HTTPResponseProperties {
	if envoy == nil {
		return nil
	}
	if config == nil {
		config = new(data.Configuration)
	}
	resp := new(md.HTTPResponseProperties)
	if envoy.GetResponseCode() != nil {
		resp.ResponseCode = envoy.GetResponseCode().Value
	}
	resp.ResponseHeadersBytes = envoy.GetResponseHeadersBytes()
	resp.ResponseBodyBytes = envoy.GetResponseBodyBytes()
	if traffic == md.Common_Traffic_Egress {
		resp.ResponseHeaders = util.MapSubset(config.Egress.ResponseHeaders, envoy.GetResponseHeaders())
		resp.ResponseTrailers = util.MapSubset(config.Egress.ResponseTrailers, envoy.GetResponseTrailers())
	} else {
		resp.ResponseHeaders = util.MapSubset(config.Ingress.ResponseHeaders, envoy.GetResponseHeaders())
		resp.ResponseTrailers = util.MapSubset(config.Ingress.ResponseTrailers, envoy.GetResponseTrailers())

	}
	resp.ResponseCodeDetails = dictionary.Add(envoy.GetResponseCodeDetails())

	return resp
}
