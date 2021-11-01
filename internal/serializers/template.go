package serializers

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"
	"unicode"
)

type TemplateContainer struct {
	Templates []string
	Data      interface{}
}

type TemplateSerializer struct{}

func NewTemplateSerializer() (Serializer, error) {
	return &TemplateSerializer{}, nil
}

func makeIndentFunc(size int) func(string) (string, error) {
	pad := strings.Repeat(" ", size)
	return func(data string) (string, error) {
		var buffer bytes.Buffer
		for _, line := range strings.Split(data, "\n") {
			fmt.Fprint(&buffer, pad)
			fmt.Fprintln(&buffer, line)
		}
		return strings.TrimRightFunc(buffer.String(), unicode.IsSpace), nil
	}
}

func (s *TemplateSerializer) Serialize(data interface{}, out io.Writer) error {
	tc, ok := data.(*TemplateContainer)
	if !ok {
		return fmt.Errorf("invalid data type passed to TemplateSerializer, must be: TemplateContainer")
	}

	rootTpl := template.New("root")

	rootTpl.Funcs(template.FuncMap{
		"indent2": makeIndentFunc(2),
		"tableify": func(data interface{}) (string, error) {
			ser, err := NewTableSerializer()
			if err != nil {
				return "", err
			}
			var buffer bytes.Buffer
			if err := ser.Serialize(data, &buffer); err != nil {
				return "", err
			}
			return buffer.String(), nil
		},
	})

	var err error
	for _, tpl := range tc.Templates {
		rootTpl, err = rootTpl.Parse(tpl)
		if err != nil {
			return err
		}
	}

	return rootTpl.Execute(out, tc.Data)
}
