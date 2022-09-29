package accesslog

import (
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/metric-data/accesslog"
	"strings"
)

func convertResponseFlags(log *accesslog.Common, envoy *datav3.AccessLogCommon) {
	var flags string

	resp := envoy.GetResponseFlags()
	if log == nil || envoy == nil || resp == nil {
		return
	}

	// Start with response code generating flags, 404.429.503.504

	// Indicates there was an upstream request timeout.
	if resp.GetUpstreamRequestTimeout() {
		flags = addFlag(flags, UPSTREAM_REQUEST_TIMEOUT)
	}

	// Indicates there was no healthy upstream.
	if resp.GetNoHealthyUpstream() {
		flags = addFlag(flags, NO_HEALTHY_UPSTREAM)
	}
	// Indicates there was a local reset by a connection pool due to an initial connection failure.
	if resp.GetUpstreamConnectionFailure() {
		flags = addFlag(flags, UPSTREAM_CONNECTION_FAILURE)
	}
	// Indicates remote codec level reset was received on the stream.
	if resp.GetUpstreamRemoteReset() {
		flags = addFlag(flags, UPSTREAM_REMOTE_RESET)
	}
	// Indicates the stream was reset due to an upstream connection termination.
	if resp.GetUpstreamConnectionTermination() {
		flags = addFlag(flags, UPSTREAM_CONNECTION_TERMINATION)
	}
	// Indicates the stream was reset because of a resource overflow.
	if resp.GetUpstreamOverflow() {
		flags = addFlag(flags, UPSTREAM_OVERFLOW)
	}
	// Indicates no route was found for the request.
	if resp.GetNoRouteFound() {
		flags = addFlag(flags, NO_ROUTE_FOUND)
	}
	// Indicates that the request was rate-limited locally.
	if resp.GetRateLimited() {
		flags = addFlag(flags, RATE_LIMITED)
	}

	// Next, finish upstream errors
	// Indicates that the upstream retry limit was exceeded, resulting in a downstream error.
	if resp.GetUpstreamRetryLimitExceeded() {
		flags = addFlag(flags, UPSTREAM_RETRY_LIMIT_EXCEEDED)
	}
	// Indicates no cluster was found for the request.
	if resp.GetNoClusterFound() {
		flags = addFlag(flags, NO_CLUSTER_FOUND)
	}

	// Next, local/downstream connection

	// Indicates local codec level reset was sent on the stream.
	if resp.GetLocalReset() {
		flags = addFlag(flags, LOCAL_RESET)
	}
	// Indicates the stream was reset due to a downstream connection termination.
	if resp.GetDownstreamConnectionTermination() {
		flags = addFlag(flags, DOWNSTREAM_CONNECTION_TERMINATION)
	}
	// Indicates local server healthcheck failed.
	if resp.GetFailedLocalHealthcheck() {
		flags = addFlag(flags, FAILED_LOCAL_HEALTH_CHECK)
	}
	// Indicates that the stream idle timeout was hit, resulting in a downstream 408.
	if resp.GetStreamIdleTimeout() {
		flags = addFlag(flags, STREAM_IDLE_TIMEOUT)
	}

	// Add UAEX

	// Indicates that the request was delayed before proxying.
	if resp.GetDelayInjected() {
		flags = addFlag(flags, DELAY_INJECTED)
	}
	// Indicates that the request was aborted with an injected error code.
	if resp.GetFaultInjected() {
		flags = addFlag(flags, FAULT_INJECTED)
	}
	// Indicates if the request was deemed unauthorized and the reason for it.
	//UnauthorizedDetails *ResponseFlags_Unauthorized
	// Indicates that the request was rejected because there was an error in rate limit service.
	if resp.GetRateLimitServiceError() {
		flags = addFlag(flags, RATELIMIT_SERVICE_ERROR)
	}
	// Indicates that the request was rejected because an envoy request header failed strict
	// validation.
	if resp.GetInvalidEnvoyRequestHeaders() {
		flags = addFlag(flags, INVALID_ENVOY_REQUEST_HEADERS)
	}
	// Indicates there was an HTTP protocol error on the downstream request.
	if resp.GetDownstreamProtocolError() {
		flags = addFlag(flags, DOWNSTREAM_PROTOCOL_ERROR)
	}
	// Indicates there was a max stream duration reached on the upstream request.
	if resp.GetUpstreamMaxStreamDurationReached() {
		flags = addFlag(flags, UPSTREAM_MAX_STREAM_DURATION_REACHED)
	}
	// Indicates the response was served from a cache filter.
	if resp.GetResponseFromCacheFilter() {
		flags = addFlag(flags, RESPONSE_FROM_CACHE_FILTER)
	}
	// Indicates that a filter configuration is not available.
	if resp.GetNoFilterConfigFound() {
		flags = addFlag(flags, NO_FILTER_CONFIG_FOUND)
	}
	// Indicates that request or connection exceeded the downstream connection duration.
	if resp.GetDurationTimeout() {
		flags = addFlag(flags, DURATION_TIMEOUT)
	}
	// Indicates there was an HTTP protocol error in the upstream response.
	if resp.GetUpstreamProtocolError() {
		flags = addFlag(flags, UPSTREAM_PROTOCOL_ERROR)
	}
	// Indicates overload manager terminated the request.
	if resp.GetOverloadManager() {
		flags = addFlag(flags, OVERLOAD_MANAGER)
	}
	// Indicates a DNS resolution failed.
	if resp.GetDnsResolutionFailure() {
		flags = addFlag(flags, DNS_FAIL)
	}
	if flags == "" {
		return
	}
	// Put first flag in ResponseFlag, the rest of the flags will be a list on the ext string
	ls := strings.Split(flags, ",")
	if len(ls) > 0 {
		log.ResponseFlag = ls[0]
		for i, s := range ls {
			if i == 0 {
				log.ResponseFlag = s
			} else {
				log.ResponseFlagExt = addFlag(log.ResponseFlagExt, s)
			}
		}
	}

}

func addFlag(flags string, token string) string {
	if flags != "" {
		flags += ","
	}
	flags += token
	return flags
}
