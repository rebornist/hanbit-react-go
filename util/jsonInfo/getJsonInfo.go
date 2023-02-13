package jsoninfo

import "encoding/json"

func (jc *JSONConfig) GetJSONInfo(name string) ([]byte, error) {
	return json.Marshal(jc.data[name])
}
