package util

import (
	"encoding/json"
	"net/http/httptest"
)

func ConvertToString(req interface{}) string {
	if req == nil {
		return ""
	}
	res, _ := json.Marshal(&req)
	return string(res)
}

func ConvertToStruct(resp *httptest.ResponseRecorder, res interface{}) error {
	return json.Unmarshal([]byte(resp.Body.Bytes()), &res)
}
