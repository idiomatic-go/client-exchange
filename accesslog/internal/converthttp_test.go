package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
	data "github.com/idiomatic-go/entity-data/accesslog"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ExampleConvertHttpAccessLogEntryPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpAccessLogEntry(md.Common_Traffic_Egress, nil, nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpAccessLogEntry(md.Common_Traffic_Egress, new(data.Configuration), util.CreateInvertedDictionary(false), new(datav3.HTTPAccessLogEntry)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}

func ExampleConvertHttpRequestPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpRequest(md.Common_Traffic_Ingress, nil, nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpRequest(md.Common_Traffic_Ingress, new(data.Configuration), util.CreateInvertedDictionary(false), new(datav3.HTTPRequestProperties)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}

func ExampleConvertHttpResponsePanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpResponse(md.Common_Traffic_Ingress, nil, nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpResponse(md.Common_Traffic_Ingress, new(data.Configuration), util.CreateInvertedDictionary(false), new(datav3.HTTPResponseProperties)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
