package util

import "encoding/json"

func ReflectToString(obj interface{}, args ...interface{}) string {
	if obj == nil {
		return "<nil>"
	}
	switch v := obj.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		result, err := json.Marshal(obj)
		if err != nil {
			return err.Error()
		} else {
			return string(result)
		}
	}

}
