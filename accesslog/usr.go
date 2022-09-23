package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"github.com/idiomatic-go/entity-data/accesslog"
)

const (
	Uri = "accesslog"
)

var View *util.VersionedEntity

func init() {
	View = accesslog.CreateVersionedEntity()
}
