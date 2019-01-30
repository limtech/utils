package utils

import (
	"fmt"
	"math"
)

// format size
func FormatSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	i := 0
	base := math.Pow(2, 10)

	x := float64(size)

	for {
		i++
		x = float64(size) / float64(math.Pow(base, float64(i)))
		if x < float64(math.Pow(base, float64(i))) {
			break
		}
	}
	return fmt.Sprintf("%.2f %s", x, units[i])
}
