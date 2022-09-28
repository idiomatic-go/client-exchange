package accesslog

import (
	"github.com/idiomatic-go/common-lib/httpxt"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	"os"
	"time"
)

var entity = data.CreateVersionedEntity()

var handler util.Niladic = func() {
	exchange()
}

var exchange util.NiladicStatus = func() error {
	var status *httpxt.ResponseStatus

	view := data.View{}
	for i := 0; i < 2; i++ {
		status := util.HttpGetContent(nil, getEntityDataUrl(), nil, &view)
		if status.IsSuccess() {
			entity.SetEntity(&view)
			return nil
		}
		if !status.IsRetriable() {
			break
		}
	}
	return status.FirstError()
}

// discoveryStartup - Initialize discovery
func discoveryStartup() {
	// Initialize entity, accessing attributes from environment if configured
	ingress := data.CSVAttributes{App: os.Getenv(LogAppIngressKey), Custom: os.Getenv(LogCustomIngressKey), RequestHeaders: os.Getenv(LogRequestHeadersIngressKey), ResponseHeaders: os.Getenv(LogResponseHeadersIngressKey), ResponseTrailers: os.Getenv(LogResponseTrailersIngressKey), Cookies: os.Getenv(LogCookiesIngressKey)}
	egress := data.CSVAttributes{App: os.Getenv(LogAppEgressKey), Custom: os.Getenv(LogCustomEgressKey), RequestHeaders: os.Getenv(LogRequestHeadersEgressKey), ResponseHeaders: os.Getenv(LogResponseHeadersEgressKey), ResponseTrailers: os.Getenv(LogResponseTrailersEgressKey), Cookies: os.Getenv(LogCookiesEgressKey)}
	view := data.CreateEntity(&ingress, &egress)
	entity.SetEntity(&view)

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
	minutes := getPollingInterval()
	if minutes > 0 {
		go util.Timer(true, time.Minute*time.Duration(minutes), nil, handler)
	}
}

func GetEntity() data.View {
	return entity.GetEntity()
}
