package serializers

import (
	"fmt"
	"io"
)

type Serializer interface {
	Serialize(interface{}, io.Writer) error
}

func NewSerializer(output string) (Serializer, error) {
	switch output {
	default:
	case "json":
		return NewJSONSerializer(), nil
	case "table":
		return NewTableSerializer()
	}
	return nil, fmt.Errorf("output: %s not recognized/supported", output)
}
