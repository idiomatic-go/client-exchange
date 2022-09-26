package accesslog

import (
	"github.com/idiomatic-go/client-exchange/accesslog/envoy"
)

func ProcessMessage(msg *accesslogv3.StreamAccessLogsMessage) {
	if msg != nil {
		// convert the message
		// send on channel
	}

}
