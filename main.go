package json

import (
	"encoding/json"
	"strconv"

	"github.com/korjavin/go-php-serialize"
	"golang.org/x/text/encoding/unicode"
)

func DecodeToJSON(val1 string) (string, error) {
	empty := ""
	val, err := phpserialize.DecodeWithEncoding(val1, unicode.UTF8)
	if err != nil {
		return empty, err
	}
	stringMap := modifyMap(val)
	jsonString, err := json.Marshal(stringMap)
	if err != nil {
		return empty, err
	}
	return string(jsonString), nil
}

func modifyMap(val interface{}) interface{} {
	stringMap := make(map[string]interface{})
	if map1, ok := val.(map[interface{}]interface{}); ok {
		for k, v := range map1 {
			stringKey := modifyValue(k)
			leaf1 := modifyMap(v)
			stringMap[stringKey] = leaf1
		}
	} else {
		return modifyValue(val)
	}
	return stringMap
}
func modifyValue(val interface{}) string {
	if s, ok := val.(string); ok {
		return s
	}
	if i, ok := val.(int64); ok {
		return strconv.FormatInt(i, 10)
	}
	return ""
}
