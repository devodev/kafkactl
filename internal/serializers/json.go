package serializers

import (
	"encoding/json"
	"io"
)

type JSONSerializer struct{}

func NewJSONSerializer() (*JSONSerializer, error) {
	return &JSONSerializer{}, nil
}

func (s *JSONSerializer) Serialize(data interface{}, out io.Writer) error {
	return json.NewEncoder(out).Encode(data)
}
