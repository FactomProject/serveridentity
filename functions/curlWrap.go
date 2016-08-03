package functions

import (
	"strings"

	"github.com/FactomProject/serveridentity/identity"
)

func CurlWrapGET(str string) string {
	return curlWrap("GET", str)
}

func CurlWrapPOST(str string) string {
	return curlWrap("POST", str)
}

func curlWrap(req string, str string) string {
	curlStr := ""
	if identity.Version == 2 {
		curlStr = "curl -X " + req + " --data '" + str + "' -H 'content-type:text/plain;' http://localhost:8088/v2"
	} else {
		if strings.Contains(str, "CommitEntryMsg") {
			curlStr = "curl -X POST --data '" + str + "' -H 'content-type:text/plain;' http://localhost:8088/v1/commit-entry"
		} else if strings.Contains(str, "CommitChainMsg") {
			curlStr = "curl -X POST --data '" + str + "' -H 'content-type:text/plain;' http://localhost:8088/v1/commit-chain"
		} else {
			curlStr = "curl -X POST --data '" + str + "' -H 'content-type:text/plain;' http://localhost:8088/v1/reveal-entry"
		}
	}
	return curlStr
}
