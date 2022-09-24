package accesslog

import (
	"github.com/idiomatic-go/common-lib/vhost"
	"os"
)

var c = make(chan vhost.Message, 10)
var startup = false

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c)
	go receive()
}

func Startup() {
	// Need to verify that there is an entity data url
	url := os.Getenv(EntityDataUrlKey)
	if url == "" {
		msg := vhost.CreateMessage(vhost.ErrorEvent, Uri, "Upstream entity data Url is empty")
		vhost.SendResponse(msg)
		return
	}
}

func Shutdown() {
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
				if !startup {
					startup = true
					Startup()
				}
			case vhost.ShutdownEvent:
				Shutdown()
			}
		}
	}
}
