package accesslog

import (
	"github.com/idiomatic-go/common-lib/httpxt"
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/common-lib/vhost"
	data "github.com/idiomatic-go/entity-data/accesslog"
	"net/http"
	"time"
)

var entity = data.CreateVersionedEntity()

var exchange util.DispatchStatus func() error {
	req,err0 := http.NewRequest("",GetEntityDataUrl(),nil)
	if err0 != nil {
		vhost.LogPrintf("%v",err0)
		return err0
	}
	status := httpxt.DoWithStatus(req)
	if status.HttpErr != nil {
		vhost.LogPrintf("%v",err0)
		return status.HttpErr
	}
}

var handler util.Dispatch = func() {
    exchange()
}

// configStartup - Initialize timer to poll for entity updates
func configStartup() {
    if exchange() != nil {
		vhost.SendErrorResponse(Uri)
		return
	}
   go util.Timer(true,time.Second * time.Duration(GetPollingInterval()),nil,handler)
}

func GetConfig() data.AccessLogView {
	return entity.GetEntity()
}


