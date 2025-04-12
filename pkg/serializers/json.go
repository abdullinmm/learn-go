package serializers

import "encoding/json"

type JSONSerializer struct{}

func (j JSONSerializer) Serializer(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
