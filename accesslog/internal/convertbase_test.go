package internal

import (
	"fmt"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

func ExampleConvertTlsPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertTls(nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertTls(new(datav3.TLSProperties)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false
}

func ExampleConvertAddressPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertAddress(nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertAddress(new(corev3.Address)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
