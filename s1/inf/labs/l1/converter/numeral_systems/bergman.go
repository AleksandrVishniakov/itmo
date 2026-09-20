package numeralsystems

import (
	"fmt"
	"math"
	"strings"
)

type BergmanNumeralSystem struct {
	base float64
}

func NewBergmanNumeralSystem() *BergmanNumeralSystem {
	return &BergmanNumeralSystem{
		base: (1.0 + math.Sqrt(5)) / 2.0,
	}
}

func (b *BergmanNumeralSystem) Name() string {
	return "Bergman"
}

func (b *BergmanNumeralSystem) Base() any {
	return b.base
}

func (b *BergmanNumeralSystem) Parse(str string) (float64, error) {
	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		dotIndex = len(str)
	}

	str = strings.Replace(str, ".", "", 1)

	var result float64

	for i, ch := range str {
		if i == dotIndex {
			continue
		}

		switch ch {
		case '0':
			continue
		case '1':
			result += math.Pow(b.base, float64(dotIndex-i-1))
		default:
			return 0, fmt.Errorf("unexpected symbol: %q", ch)
		}
	}

	return result, nil
}

func (b *BergmanNumeralSystem) Format(n float64) (string, error) {
	var result string

	i := 0
	for ; math.Pow(b.base, float64(i)) < n; i++ {
	}
	if math.Pow(b.base, float64(i)) > n {
		i -= 1
	}

	for ; n > 0.01; i-- {
		if v := math.Pow(b.base, float64(i)); v <= n {
			n -= v
			result += "1"
		} else {
			result += "0"
		}

		if i == 0 && n != 0 {
			result += "."
		}
	}

	return result, nil
}
