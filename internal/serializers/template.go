package serializers

import (
	"fmt"
	"io"
	"text/template"
)

type TemplateContainer struct {
	Templates []string
	Data      interface{}
}

type TemplateSerializer struct{}

func NewTemplateSerializer() (Serializer, error) {
	return &TemplateSerializer{}, nil
}

func (s *TemplateSerializer) Serialize(data interface{}, out io.Writer) error {
	tc, ok := data.(*TemplateContainer)
	if !ok {
		return fmt.Errorf("invalid data type passed to TemplateSerializer, must be: TemplateContainer")
	}

	rootTpl := template.New("root")

	var err error
	for _, tpl := range tc.Templates {
		rootTpl, err = rootTpl.Parse(tpl)
		if err != nil {
			return err
		}
	}

	return rootTpl.Execute(out, tc.Data)
}
