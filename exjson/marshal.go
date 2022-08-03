package exjson

import (
	"bytes"
	"encoding/json"
)

func MarshalIndent(data any, indent string) ([]byte, error) {
	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if indent != "" {
		buf := new(bytes.Buffer)
		if err = json.Indent(buf, result, "", indent); err == nil {
			return buf.Bytes(), nil
		}
	}

	return result, nil
}
