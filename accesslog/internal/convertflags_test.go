package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

func ExampleConvertResponseFlagsPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertResponseFlags(nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertResponseFlags(new(datav3.ResponseFlags)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
