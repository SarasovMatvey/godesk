package helper

import (
	"encoding/json"
	"io"
	"strings"
)

func ObjectToJsonReader(obj interface{}) io.Reader {
	bytes, _ := json.Marshal(obj)

	return strings.NewReader(string(bytes))
}
