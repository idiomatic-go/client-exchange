package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"os"
	"strconv"
)

const (
	// Environment keys for access logging mutations
	AccessLogIngress          = "ACCESS_LOG_INGRESS"
	AccessLogEgress           = "ACCESS_LOG_EGRESS"
	AccessLogRequestHeaders   = "ACCESS_LOG_REQUEST_HEADERS"
	AccessLogResponseHeaders  = "ACCESS_LOG_RESPONSE_HEADERS"
	AccessLogResponseTrailers = "ACCESS_LOG_RESPONSE_TRAILERS"
	AccessLogCookies          = "ACCESS_LOG_COOKIES"

	// PollingIntervalKey - Environment key for configuration polling interval
	AccessLogPollingIntervalKey = "ACCESS_LOG_POLLING_INTERVAL" // Numeric string denoting minutes

	// EntityDataUrlKey - Environment key for entity data server Url
	AccessLogEntityDataUrlKey = "ACCESS_LOG_ENTITY_DATA_URL"
)

func getPollingInterval() (minutes int) {
	s := os.Getenv(AccessLogPollingIntervalKey)
	if s == "" {
		return 0
	}
	minutes, err := strconv.Atoi(s)
	if err != nil {
		util.LogPrintf("%v", err)
		return 0
	}
	if minutes < 0 {
		return 0
	}
	return minutes
}

func getEntityDataUrl() string {
	return os.Getenv(AccessLogEntityDataUrlKey)
}
