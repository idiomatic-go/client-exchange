package accesslog

import (
	"github.com/idiomatic-go/client-exchange/accesslog/envoy"
	mdata "github.com/idiomatic-go/metric-data/accesslog"
)

func ProcessMessage(msg *accesslogv3.StreamAccessLogsMessage) {
	if msg != nil {
		// convert the message
		// send on channel
	}

}

func convert(msg *accesslogv3.StreamAccessLogsMessage) mdata.StreamAccessLogsMessage {
	var d mdata.StreamAccessLogsMessage

	return d
}
