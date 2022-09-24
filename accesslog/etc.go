package accesslog

import (
	"github.com/idiomatic-go/common-lib/vhost"
	"os"
	"strconv"
)

const (
	// PollingIntervalKey - Environment key for configuration polling interval
	AccessLogPollingIntervalKey = "ACCESS_POLLING_INTERVAL" // Numeric string denoting seconds

	DefaultPollingInterval = 600 // 10 minutes

	// EntityDataUrlKey - Environment key for entity data server Url
	EntityDataUrlKey = "ENTITY_DATA_URL"
)

func GetPollingInterval() (seconds int) {
	seconds = DefaultPollingInterval
	s := os.Getenv(AccessLogPollingIntervalKey)
	if s != "" {
		sec, err := strconv.Atoi(s)
		if err != nil {
			vhost.LogPrintf("%v", err)
		}
		seconds = sec
	}
	return seconds
}

func GetEntityDataUrl() string {
	return os.Getenv(EntityDataUrlKey)
}
