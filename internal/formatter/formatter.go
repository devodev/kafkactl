package formatter

import (
	"regexp"
	"strings"
)

var (
	camelCaseRe = regexp.MustCompile(`^[a-z]|_[a-z]`)
)

type Header func(header string) string

func SnakeCaseToUpperCase(header string) string {
	header = strings.ReplaceAll(header, "_", "-")
	header = strings.ToUpper(header)
	return header
}

func SnakeCaseToCamelCase(header string) string {
	header = strings.ReplaceAll(header, "_", "-")
	header = camelCaseRe.ReplaceAllStringFunc(header, func(w string) string {
		return strings.ToUpper(w)
	})
	return header
}

type RowValue func(header, value string, width int) string

func EmptyPlaceholder(placeholder string) RowValue {
	return func(header, value string, width int) string {
		if value == "" {
			value = placeholder
		}
		valueR := []rune(value)
		if len(valueR) > width {
			valueR = valueR[:width]
		}
		return string(valueR)
	}
}
