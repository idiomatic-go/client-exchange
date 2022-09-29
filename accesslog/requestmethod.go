package accesslog

// Note : had to create this from github, as the go-control-plane module does not include
// this source

// RequestMethod - HTTP request method.
type RequestMethod int32

const (
	RequestMethod_METHOD_UNSPECIFIED RequestMethod = 0
	RequestMethod_GET                RequestMethod = 1
	RequestMethod_HEAD               RequestMethod = 2
	RequestMethod_POST               RequestMethod = 3
	RequestMethod_PUT                RequestMethod = 4
	RequestMethod_DELETE             RequestMethod = 5
	RequestMethod_CONNECT            RequestMethod = 6
	RequestMethod_OPTIONS            RequestMethod = 7
	RequestMethod_TRACE              RequestMethod = 8
	RequestMethod_PATCH              RequestMethod = 9
)

// Enum value maps for RequestMethod.
var (
	RequestMethod_name = map[int32]string{
		0: "METHOD_UNSPECIFIED",
		1: "GET",
		2: "HEAD",
		3: "POST",
		4: "PUT",
		5: "DELETE",
		6: "CONNECT",
		7: "OPTIONS",
		8: "TRACE",
		9: "PATCH",
	}
	RequestMethod_value = map[string]int32{
		"METHOD_UNSPECIFIED": 0,
		"GET":                1,
		"HEAD":               2,
		"POST":               3,
		"PUT":                4,
		"DELETE":             5,
		"CONNECT":            6,
		"OPTIONS":            7,
		"TRACE":              8,
		"PATCH":              9,
	}
)

func GetRequestMethodName(method int32) string {
	if method < 0 || method > int32(RequestMethod_PATCH) {
		return ""
	}
	return RequestMethod_name[method]
}
