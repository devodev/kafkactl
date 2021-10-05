package presentation

import (
	"strconv"
	"strings"
)

func intSliceToString(a []int, delim string) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, delim)
}
