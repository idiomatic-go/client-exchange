package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	"os"
	"time"
)

var entity = data.CreateVersionedEntity()

var exchange util.DispatchStatus = func() error {
	mu := data.Mutations{}
	status := util.HttpGetContent(nil, getEntityDataUrl(), nil, &mu)
	if status.IsSuccess() {
		entity.SetEntity(&mu)
	}
	return status.FirstError()
}

var handler util.Dispatch = func() {
	exchange()
}

// discoveryStartup - Initialize discovery
func discoveryStartup() {
	// Initialize entity, accessing attributes from environment if configured
	mu := data.CreateEntity(os.Getenv(AccessLogIngress), os.Getenv(AccessLogEgress), os.Getenv(AccessLogRequestHeaders), os.Getenv(AccessLogResponseHeaders), os.Getenv(AccessLogResponseTrailers), os.Getenv(AccessLogCookies))
	entity.SetEntity(&mu)

	// If no Url configured, then discovery is disabled
	uri := getEntityDataUrl()
	if uri == "" {
		return
	}
	
	// Discovery is enabled, so do an initial exchange. If there is an error, then do not configure polling
	err := exchange()
	if err != nil {
		return
	}

	// Check the polling interval to see if polling is enabled.
	seconds := getPollingInterval()
	if seconds > 0 {
		go util.Timer(true, time.Minute*time.Duration(getPollingInterval()), nil, handler)
	}
}

func GetEntity() data.Mutations {
	return entity.GetEntity()
}
