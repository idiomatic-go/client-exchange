package accesslog

import (
	"github.com/idiomatic-go/common-lib/vhost"
	vusr "github.com/idiomatic-go/common-lib/vhost/usr"
)

var c = make(chan *vusr.Message, 10)
var startup = false

// init - registers package with a channel
func init() {
	vhost.RegisterPackage(Uri, c)
	go receive()
}

func Startup() {

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
			case vusr.StartupEvent:
				if !startup {
					startup = true
					Startup()
				}
			case vusr.ShutdownEvent:
				Shutdown()
			}
		}
	}
}
