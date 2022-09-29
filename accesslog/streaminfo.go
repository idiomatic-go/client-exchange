package accesslog

const (
	NONE                                 = "-"
	DOWNSTREAM_CONNECTION_TERMINATION    = "DC"
	FAILED_LOCAL_HEALTH_CHECK            = "LH"
	NO_HEALTHY_UPSTREAM                  = "UH"
	UPSTREAM_REQUEST_TIMEOUT             = "UT"
	LOCAL_RESET                          = "LR"
	UPSTREAM_REMOTE_RESET                = "UR"
	UPSTREAM_CONNECTION_FAILURE          = "UF"
	UPSTREAM_CONNECTION_TERMINATION      = "UC"
	UPSTREAM_OVERFLOW                    = "UO"
	UPSTREAM_RETRY_LIMIT_EXCEEDED        = "URX"
	NO_ROUTE_FOUND                       = "NR"
	DELAY_INJECTED                       = "DI"
	FAULT_INJECTED                       = "FI"
	RATE_LIMITED                         = "RL"
	UNAUTHORIZED_EXTERNAL_SERVICE        = "UAEX"
	RATELIMIT_SERVICE_ERROR              = "RLSE"
	STREAM_IDLE_TIMEOUT                  = "SI"
	INVALID_ENVOY_REQUEST_HEADERS        = "IH"
	DOWNSTREAM_PROTOCOL_ERROR            = "DPE"
	UPSTREAM_MAX_STREAM_DURATION_REACHED = "UMSDR"
	RESPONSE_FROM_CACHE_FILTER           = "RFCF"
	NO_FILTER_CONFIG_FOUND               = "NFCF"
	DURATION_TIMEOUT                     = "DT"
	UPSTREAM_PROTOCOL_ERROR              = "UPE"
	NO_CLUSTER_FOUND                     = "NC"
	OVERLOAD_MANAGER                     = "OM"
	DNS_FAIL                             = "DF"
)

/*
52   static constexpr std::array ALL_RESPONSE_STRING_FLAGS{
53       FlagStringAndEnum{FAILED_LOCAL_HEALTH_CHECK, ResponseFlag::FailedLocalHealthCheck},
54       FlagStringAndEnum{NO_HEALTHY_UPSTREAM, ResponseFlag::NoHealthyUpstream},
55       FlagStringAndEnum{UPSTREAM_REQUEST_TIMEOUT, ResponseFlag::UpstreamRequestTimeout},
56       FlagStringAndEnum{LOCAL_RESET, ResponseFlag::LocalReset},
57       FlagStringAndEnum{UPSTREAM_REMOTE_RESET, ResponseFlag::UpstreamRemoteReset},
58       FlagStringAndEnum{UPSTREAM_CONNECTION_FAILURE, ResponseFlag::UpstreamConnectionFailure},
59       FlagStringAndEnum{UPSTREAM_CONNECTION_TERMINATION,
60                         ResponseFlag::UpstreamConnectionTermination},
61       FlagStringAndEnum{UPSTREAM_OVERFLOW, ResponseFlag::UpstreamOverflow},
62       FlagStringAndEnum{NO_ROUTE_FOUND, ResponseFlag::NoRouteFound},
63       FlagStringAndEnum{DELAY_INJECTED, ResponseFlag::DelayInjected},
64       FlagStringAndEnum{FAULT_INJECTED, ResponseFlag::FaultInjected},
65       FlagStringAndEnum{RATE_LIMITED, ResponseFlag::RateLimited},
66       FlagStringAndEnum{UNAUTHORIZED_EXTERNAL_SERVICE, ResponseFlag::UnauthorizedExternalService},
67       FlagStringAndEnum{RATELIMIT_SERVICE_ERROR, ResponseFlag::RateLimitServiceError},
68       FlagStringAndEnum{DOWNSTREAM_CONNECTION_TERMINATION,
69                         ResponseFlag::DownstreamConnectionTermination},
70       FlagStringAndEnum{UPSTREAM_RETRY_LIMIT_EXCEEDED, ResponseFlag::UpstreamRetryLimitExceeded},
71       FlagStringAndEnum{STREAM_IDLE_TIMEOUT, ResponseFlag::StreamIdleTimeout},
72       FlagStringAndEnum{INVALID_ENVOY_REQUEST_HEADERS, ResponseFlag::InvalidEnvoyRequestHeaders},
73       FlagStringAndEnum{DOWNSTREAM_PROTOCOL_ERROR, ResponseFlag::DownstreamProtocolError},
74       FlagStringAndEnum{UPSTREAM_MAX_STREAM_DURATION_REACHED,
75                         ResponseFlag::UpstreamMaxStreamDurationReached},
76       FlagStringAndEnum{RESPONSE_FROM_CACHE_FILTER, ResponseFlag::ResponseFromCacheFilter},
77       FlagStringAndEnum{NO_FILTER_CONFIG_FOUND, ResponseFlag::NoFilterConfigFound},
78       FlagStringAndEnum{DURATION_TIMEOUT, ResponseFlag::DurationTimeout},
79       FlagStringAndEnum{UPSTREAM_PROTOCOL_ERROR, ResponseFlag::UpstreamProtocolError},
80       FlagStringAndEnum{NO_CLUSTER_FOUND, ResponseFlag::NoClusterFound},
81       FlagStringAndEnum{OVERLOAD_MANAGER, ResponseFlag::OverloadManager},
82   };
83

*/

const (
	// Local server healthcheck failed.
	FailedLocalHealthCheck = 0x1
	// No healthy upstream.
	NoHealthyUpstream = 0x2
	// Request timeout on upstream.
	UpstreamRequestTimeout = 0x4
	// Local codec level reset was sent on the stream.
	LocalReset = 0x8
	// Remote codec level reset was received on the stream.
	UpstreamRemoteReset = 0x10
	// Local reset by a connection pool due to an initial connection failure.
	UpstreamConnectionFailure = 0x20
	// If the stream was locally reset due to connection termination.
	UpstreamConnectionTermination = 0x40
	// The stream was reset because of a resource overflow.
	UpstreamOverflow = 0x80
	// No route found for a given request.
	NoRouteFound = 0x100
	// Request was delayed before proxying.
	DelayInjected = 0x200
	// Abort with error code was injected.
	FaultInjected = 0x400
	// Request was ratelimited locally by rate limit filter.
	RateLimited = 0x800
	// Request was unauthorized by external authorization service.
	UnauthorizedExternalService = 0x1000
	// Unable to call Ratelimit service.
	RateLimitServiceError = 0x2000
	// If the stream was reset due to a downstream connection termination.
	DownstreamConnectionTermination = 0x4000
	// Exceeded upstream retry limit.
	UpstreamRetryLimitExceeded = 0x8000
	// Request hit the stream idle timeout, triggering a 408.
	StreamIdleTimeout = 0x10000
	// Request specified x-envoy-* header values that failed strict header checks.
	InvalidEnvoyRequestHeaders = 0x20000
	// Downstream request had an HTTP protocol error
	DownstreamProtocolError = 0x40000
	// Upstream request reached to user defined max stream duration.
	UpstreamMaxStreamDurationReached = 0x80000
	// True if the response was served from an Envoy cache filter.
	ResponseFromCacheFilter = 0x100000
	// Filter config was not received within the permitted warming deadline.
	NoFilterConfigFound = 0x200000
	// Request or connection exceeded the downstream connection duration.
	DurationTimeout = 0x400000
	// Upstream response had an HTTP protocol error
	UpstreamProtocolError = 0x800000
	// No cluster found for a given request.
	NoClusterFound = 0x1000000
	// Overload Manager terminated the stream.
	OverloadManager = 0x2000000
	// DNS resolution failed.
	DnsResolutionFailed = 0x4000000
	// ATTENTION: MAKE SURE THIS REMAINS EQUAL TO THE LAST FLAG.
	LastFlag = DnsResolutionFailed
)
