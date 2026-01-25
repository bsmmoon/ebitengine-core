package tiles

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed layers/background.csv
var backgroundLayerData string

//go:embed layers/foreground.csv
var foregroundLayerData string

// parseCSVLayer parses a CSV string into a slice of integers.
func parseCSVLayer(data string) []int {
	var result []int
	lines := strings.Split(strings.TrimSpace(data), "\n")
	for _, line := range lines {
		values := strings.Split(line, ",")
		for _, v := range values {
			num, err := strconv.Atoi(strings.TrimSpace(v))
			if err == nil {
				result = append(result, num)
			}
		}
	}
	return result
}
