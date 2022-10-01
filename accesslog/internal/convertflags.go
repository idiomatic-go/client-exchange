package internal

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ConvertResponseFlags(envoy *datav3.ResponseFlags) *md.ResponseFlags {
	if envoy == nil {
		return nil
	}
	resp := new(md.ResponseFlags)
	resp.Encoded = EncodeResponseFlags(envoy)
	if envoy.GetUnauthorizedDetails() != nil {
		resp.UnauthorizedDetails = new(md.ResponseFlags_Unauthorized)
		resp.UnauthorizedDetails.Reason = md.ResponseFlags_Unauthorized_Reason(envoy.GetUnauthorizedDetails().Reason)
	}
	return resp
}
