package serializers

import (
	"fmt"
	"io"
)

type Type int

func (t Type) String() string {
	return TypesRev[t]
}

const (
	JSON Type = iota
	Table
	Template
)

var (
	Types = map[string]Type{
		"json":     JSON,
		"table":    Table,
		"template": Template,
	}
	TypesRev = map[Type]string{
		JSON:     "json",
		Table:    "table",
		Template: "template",
	}
)

func TypeFrom(s string) (Type, error) {
	val, ok := Types[s]
	if !ok {
		return 0, fmt.Errorf("invalid type: %s", s)
	}
	return val, nil
}

type Serializer interface {
	Serialize(interface{}, io.Writer) error
}

func NewSerializer(serType string) (Serializer, error) {
	t, err := TypeFrom(serType)
	if err != nil {
		return nil, err
	}

	var ser Serializer
	switch t {
	default:
		return nil, fmt.Errorf("no serializer found for: %s", t.String())
	case JSON:
		ser, err = NewJSONSerializer()
	case Table:
		ser, err = NewTableSerializer()
	case Template:
		ser, err = NewTemplateSerializer()
	}
	return ser, err
}

func IsSupported(o string) bool {
	_, err := TypeFrom(o)
	return err == nil
}
