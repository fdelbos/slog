package slog

import (
	"fmt"
	"strings"
)

const (
	printTerminalCols = 80
)

func padding(name, value string) string {
	size := 30 - len(name)
	return strings.Repeat(".", size)
}

func line(name, value string) {
	if len(name) > 26 {
		name = name[:26]
	}
	if len(value) > 46 {
		value = value[:46]
	}

	fmt.Printf(
		"* %s%s%s\n",
		name,
		padding(name, value),
		value,
	)
}

func Service(name string) {
	if len(name) > 74 {
		name = name[:74]
	}
	fmt.Printf("\n  %s \n", strings.ToUpper(name))

}

func Value(name string, value interface{}) {
	if value == nil {
		line(name, "(nil)")
	} else {
		switch t := value.(type) {
		case string:
			if t == "" {
				line(name, "(empty)")
			} else {
				line(name, t)
			}
		case bool:
			if t {
				line(name, "TRUE")
			} else {
				line(name, "FALSE")
			}
		case uint, uint8, uint16, uint32, uint64:
			line(name, fmt.Sprintf("%d", t))
		case int, int8, int16, int32, int64:
			line(name, fmt.Sprintf("%d", t))
		case float32, float64:
			line(name, fmt.Sprintf("%f", t))
		default:
			line(name, "(present)")
		}
	}
}

func Secret(name string) {
	line(name, "(secret)")
}
