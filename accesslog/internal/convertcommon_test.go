package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ExampleConvertCommonPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertCommon(md.Common_Traffic_Ingress, nil, nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertCommon(md.Common_Traffic_Ingress, new(data.Configuration), util.CreateInvertedDictionary(false), new(datav3.AccessLogCommon)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
