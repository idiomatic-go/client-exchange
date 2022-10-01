package accesslog

import md "github.com/idiomatic-go/metric-data/accesslogv3"

func CreateMessage() *md.Message {
	msg := new(md.Message)
	//msg.Origin = createOriginIdentifier()
	return msg
}
