package accesslog

import (
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
)

func ProcessMessage(msg *servicev3.StreamAccessLogsMessage) {
	if msg != nil {
		// internal the message
		// send on channel
	}

}
