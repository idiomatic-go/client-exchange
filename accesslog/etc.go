package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
	"os"
	"strconv"
)

const (
	// Environment keys for logging views
	LogAppIngressKey = "LOG_APP_INGRESS"
	LogAppEgressKey  = "LOG_APP_EGRESS"

	LogCustomIngressKey = "LOG_CUSTOM_INGRESS"
	LogCustomEgressKey  = "LOG_CUSTOM_EGRESS"

	LogRequestHeadersIngressKey = "LOG_REQUEST_HEADERS_INGRESS"
	LogRequestHeadersEgressKey  = "LOG_REQUEST_HEADERS_EGRESS"

	LogResponseHeadersIngressKey = "LOG_RESPONSE_HEADERS_INGRESS"
	LogResponseHeadersEgressKey  = "LOG_RESPONSE_HEADERS_EGRESS"

	LogResponseTrailersIngressKey = "LOG_RESPONSE_TRAILERS_INGRESS"
	LogResponseTrailersEgressKey  = "LOG_RESPONSE_TRAILERS_EGRESS"

	LogCookiesIngressKey = "LOG_COOKIES_INGRESS"
	LogCookiesEgressKey  = "LOG_COOKIES_EGRESS"

	LogPollingIntervalKey = "LOG_POLLING_INTERVAL" // Numeric string denoting minutes
	LogEntityDataUrlKey   = "LOG_ENTITY_DATA_URL"

	LogIngressOriginHttpUrlKey     = "LOG_INGRESS_ORIGIN_HTTP_URL"
	LogIngressOriginGRPCAddressKey = "LOG_INGRESS_ORIGIN_GRPC_ADDRESS"

	IdentifierRegionKey      = "ID_REGION"
	IdentifierZoneKey        = "ID_ZONE"
	IdentifierSubZoneKey     = "ID_SUB_ZONE"
	IdentifierApplicationKey = "ID_APPLICATION"
	IdentifierInstanceId     = "ID_INSTANCE_ID"
)

func getPollingInterval() (minutes int) {
	s := os.Getenv(LogPollingIntervalKey)
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
	return os.Getenv(LogEntityDataUrlKey)
}

func getIngressHttpUrl() string {
	return os.Getenv(LogIngressOriginHttpUrlKey)
}

func getIngressGRPCAddress() string {
	return os.Getenv(LogIngressOriginGRPCAddressKey)
}

func createIdentifier() *md.Provenance {
	id := new(md.Provenance)
	id.Locality.Region = os.Getenv(IdentifierRegionKey)
	id.Locality.Zone = os.Getenv(IdentifierZoneKey)
	id.Locality.SubZone = os.Getenv(IdentifierZoneKey)
	id.Application = os.Getenv(IdentifierApplicationKey)
	id.InstanceId = os.Getenv(IdentifierInstanceId)
	return id
}
