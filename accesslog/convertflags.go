package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func convertResponseFlags(log *md.AccessLogCommon, resp *datav3.ResponseFlags) {
	if log == nil || resp == nil {
		return
	}

	log.ResponseFlags = new(md.ResponseFlags)
	log.ResponseFlags.Encoded = EncodeResponseFlags(resp)
	if resp.GetUnauthorizedDetails() != nil {
		log.ResponseFlags.UnauthorizedDetails = new(md.ResponseFlags_Unauthorized)
		log.ResponseFlags.UnauthorizedDetails.Reason = md.ResponseFlags_Unauthorized_Reason(resp.GetUnauthorizedDetails().Reason)
	}
}
