package internal

import (
	"fmt"
	servicev3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	md "github.com/idiomatic-go/metric-data/accesslogv3"
)

func ExampleConvertMessagePanicCheck() {
	ConvertMessage(nil, nil)
	fmt.Println("Panic input : false")
	ConvertMessage(new(md.Message), new(servicev3.StreamAccessLogsMessage))
	fmt.Println("Panic result : false")

	//Output:
	// Panic input : false
	// Panic result : false

}
