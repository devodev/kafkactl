package serializers

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/devodev/kafkactl/internal/formatter"
)

var (
	defaultPadding            = 3
	defaultColumnFormatter    = formatter.SnakeCaseToUpperCase
	defaultRowValueFormatters = []formatter.RowValue{formatter.EmptyPlaceholder("-")}
)

type Tableizer interface {
	TableHeader() []string
	TableRows() []map[string]string
}

type TableSerializerOption func(*TableSerializer) error

func WithPadding(p int) TableSerializerOption {
	return func(b *TableSerializer) error {
		if p < 0 {
			return fmt.Errorf("padding must be greater than or equal to 0")
		}
		b.padding = p
		return nil
	}
}

func WithColumnFormatter(f formatter.Header) TableSerializerOption {
	return func(b *TableSerializer) error {
		b.columnFormatter = f
		return nil
	}
}

func WithRowValueFormatters(f []formatter.RowValue) TableSerializerOption {
	return func(b *TableSerializer) error {
		b.rowValueFormatters = f
		return nil
	}
}

type TableSerializer struct {
	strings.Builder

	padding            int
	columnFormatter    formatter.Header
	rowValueFormatters []formatter.RowValue

	widthMap map[string]int
}

func NewTableSerializer(options ...TableSerializerOption) (*TableSerializer, error) {
	builder := &TableSerializer{
		padding:            defaultPadding,
		columnFormatter:    defaultColumnFormatter,
		rowValueFormatters: defaultRowValueFormatters,
	}
	for _, opt := range options {
		if err := opt(builder); err != nil {
			return nil, err
		}
	}
	return builder, nil
}

func (s *TableSerializer) Serialize(data interface{}, out io.Writer) error {
	t, ok := data.(Tableizer)
	if !ok {
		return fmt.Errorf("unsupported data type passed to TableSerializer.Serialize")
	}
	return s.serialize(t, out)
}

func (s *TableSerializer) serialize(data Tableizer, out io.Writer) error {
	s.Build(data.TableHeader(), data.TableRows())
	if _, err := fmt.Fprint(out, s.String()); err != nil {
		return fmt.Errorf("could not serialize: %w", err)
	}
	return nil
}

func (b *TableSerializer) Build(header []string, rows []map[string]string) {
	b.buildWidthMap(header, rows)
	b.writeHeader(header)
	b.writeRows(header, rows)
}

func (b *TableSerializer) buildWidthMap(header []string, rows []map[string]string) {
	b.widthMap = make(map[string]int, len(rows))

	for _, row := range rows {
		for _, columnName := range header {
			if value, ok := row[columnName]; ok {
				oldWidth := utf8.RuneCountInString(columnName)
				if w, ok := b.widthMap[columnName]; ok {
					oldWidth = w
				}
				b.widthMap[columnName] = max(oldWidth, utf8.RuneCountInString(value))
			}
		}
	}
}

func (b *TableSerializer) writeHeader(header []string) {
	for idx, columnName := range header {
		width := b.widthMap[columnName]
		// write column
		fmt.Fprintf(b, "%-*v", width, b.columnFormatter(columnName))
		// pad column
		if idx != len(header)-1 {
			fmt.Fprintf(b, "%s", strings.Repeat(" ", b.padding))
		}
	}
	fmt.Fprintf(b, "\n")
}

func (b *TableSerializer) writeRows(header []string, rows []map[string]string) {
	for _, row := range rows {
		addNewLine := false
		for _, columnName := range header {
			width := b.widthMap[columnName]
			if value, ok := row[columnName]; ok {
				for _, rvFormatter := range b.rowValueFormatters {
					value = rvFormatter(columnName, value, width)
				}
				fmt.Fprintf(b, "%-*v", width, value)
				fmt.Fprintf(b, "%s", strings.Repeat(" ", b.padding))
				addNewLine = true
			}
		}
		if addNewLine {
			fmt.Fprint(b, "\n")
		}
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
