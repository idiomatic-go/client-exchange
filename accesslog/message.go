package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func CreateMessage() *md.Message {
	msg := new(md.Message)
	msg.Dictionary = util.CreateInvertedDictionary(false)
	msg.Identifier = createIdentifier()
	return msg
}
