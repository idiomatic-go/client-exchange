package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
)

func ExampleConvertHttpAccessLogEntryPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpAccessLogEntry(nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpAccessLogEntry(util.CreateInvertedDictionary(false), new(datav3.HTTPAccessLogEntry)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}

func ExampleConvertHttpRequestPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpRequest(nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpRequest(util.CreateInvertedDictionary(false), new(datav3.HTTPRequestProperties)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}

func ExampleConvertHttpResponsePanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertHttpResponse(nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertHttpResponse(util.CreateInvertedDictionary(false), new(datav3.HTTPResponseProperties)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
