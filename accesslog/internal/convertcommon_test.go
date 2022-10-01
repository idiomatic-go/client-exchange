package internal

import (
	"fmt"
	datav3 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"github.com/idiomatic-go/common-lib/util"
)

func ExampleConvertCommonPanicCheck() {
	fmt.Printf("Panic input : %v\n", ConvertCommon(nil, nil) != nil)
	fmt.Printf("Panic result : %v\n", ConvertCommon(util.CreateInvertedDictionary(false), new(datav3.AccessLogCommon)) == nil)

	//Output:
	// Panic input : false
	// Panic result : false

}
