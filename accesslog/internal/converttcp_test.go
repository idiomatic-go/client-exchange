package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

func ExampleConvertTcpAccessLogEntryPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertTcpAccessLogEntry(nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertTcpAccessLogEntry(new(datav3.TCPAccessLogEntry)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
