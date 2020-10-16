package empty_status_json

import "encoding/json"

type statusCodeJson struct {
	statusCode int
}

func JsonWithStatusCode(statusCode int) (res []byte) {
	res, _ = json.Marshal(statusCodeJson{statusCode: statusCode})
	return res
}
