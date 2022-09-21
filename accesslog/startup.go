package accesslog

import (
	"github.com/idiomatic-go/entity-data/accesslog"
	"github.com/idiomatic-go/metrics-data/accesslog/timeseries/v1"
)

func Startup() {
	view := accesslog.AccessLogView{}
	if len(view.Headers) > 0 {

	}
	log := timeseries.StreamAccessLogsMessage{}
	//log := v1.StreamAccessLogsMessage{}
	if log.ID != "" {

	}
}
