package accesslog

import (
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/idiomatic-go/client-exchange/accesslog/internal"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

// Send - send the message to the upstream origin server via gRPC, with optional receive
func Send(msg *md.Message, ack bool) {
	if msg == nil {
		// Log error message
		return
	}
	if ack {

	}
}

// SendViaHttp - send the message to the upstream origin server via Http
func SendViaHttp(msg *md.Message) {
	if msg == nil {
		// Log error message
		return
	}
}

// SendEnvoyMessage - send the envoy message to the upstream origin server via gRPC
func SendEnvoyMessage(envoy *servicev3.StreamAccessLogsMessage, ack bool) {
	if envoy == nil {
		// Log error message
		return
	}
	msg := CreateMessage()
	internal.ConvertMessage(GetConfiguration(), msg, envoy)
	Send(msg, ack)
}

// SendEnvoyMessageViaHttp - send the envoy message to the upstream origin server via Http
func SendEnvoyMessageViaHttp(envoy *servicev3.StreamAccessLogsMessage) {
	if envoy == nil {
		// Log error message
		return
	}
	msg := CreateMessage()
	internal.ConvertMessage(GetConfiguration(), msg, envoy)
	SendViaHttp(msg)
}
