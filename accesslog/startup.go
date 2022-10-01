package accesslog

import (
	"github.com/idiomatic-go/common-lib/vhost"
)

var c = make(chan vhost.Message, 10)
var started = false

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c)
	go receive()
}

func startup() {
	// Package configuration validation
	// origin identifier validation
	// How to verify origin connectivity?
	// Need to determine security for upstream origin discovery requests
	discoveryStartup()

}

func shutdown() {
	vhost.UnregisterPackage(Uri)
}

func receive() {
	for {
		select {
		case msg, open := <-c:
			// Exit on a closed channel
			if !open {
				return
			}
			switch msg.Event {
			case vhost.StartupEvent:
				if !started {
					started = true
					startup()
				}
			case vhost.ShutdownEvent:
				shutdown()
			}
		}
	}
}
