package functions

func CurlWrapGET(str string) string {
	return curlWrap("GET", str)
}

func CurlWrapPOST(str string) string {
	return curlWrap("POST", str)
}

func curlWrap(req string, str string) string {
	curlStr := "curl -X " + req + " --data '" + str + "' -H 'content-type:text/plain;' http://localhost:8088/v2"
	return curlStr
}
