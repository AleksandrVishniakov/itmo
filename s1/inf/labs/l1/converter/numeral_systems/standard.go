package numeralsystems

import (
	"fmt"
	"math"
	"strings"
)

var (
	symbols    = "0123456789abcdefghijklmnopqrstuvwxyz"
	charValues = map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
		'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14,
		'f': 15, 'g': 16, 'h': 17, 'i': 18, 'j': 19,
		'k': 20, 'l': 21, 'm': 22, 'n': 23, 'o': 24,
		'p': 25, 'q': 26, 'r': 27, 's': 28, 't': 29,
		'u': 30, 'v': 31, 'w': 32, 'x': 33, 'y': 34,
		'z': 35,
	}
)

type StandardNumeralSystem struct {
	base      int
	precision int
}

func NewStandardNumeralSystem(base int) *StandardNumeralSystem {
	return &StandardNumeralSystem{
		base:      base,
		precision: 8,
	}
}

func (s *StandardNumeralSystem) Name() string {
	return fmt.Sprintf("Standard{%d}", s.base)
}

func (s *StandardNumeralSystem) Base() any {
	return s.base
}

func (s *StandardNumeralSystem) Parse(str string) (float64, error) {
	if s.base < 2 || s.base > 36 {
		return 0, fmt.Errorf("invalid base: %d", s.base)
	}

	isNegative := len(str) > 0 && str[0] == '-'
	if isNegative {
		str = str[1:]
	}

	str = strings.ToLower(str)

	values := make([]int, 0)
	dotIndex := strings.Index(str, ".")

	for i, ch := range str {
		if i == dotIndex {
			continue
		}

		if v, ok := charValues[ch]; ok {
			if v >= s.base {
				return 0, fmt.Errorf("invalid symbol %q in base %d", ch, s.base)
			}

			values = append(values, v)
		} else {
			return 0, fmt.Errorf("unexpected symbol: %q", ch)
		}
	}

	if dotIndex == -1 {
		dotIndex = len(values)
	}

	var result float64

	for i, v := range values {
		result += float64(v) * math.Pow(float64(s.base), float64(dotIndex-i-1))
	}

	if isNegative {
		result = -result
	}

	return result, nil
}

func (s *StandardNumeralSystem) Format(n float64) (string, error) {
	if s.base < 2 || s.base > 36 {
		return "", fmt.Errorf("invalid base: %d", s.base)
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	var result string

	integerPart := int(n)

	if integerPart == 0 {
		result = "0"
	}

	for integerPart > 0 {
		result = string(symbols[integerPart%s.base]) + result
		integerPart /= s.base
	}

	_, fractionalPart := math.Modf(n)
	if fractionalPart > 0 {
		result += "."
	}

	for i := 0; i < s.precision && fractionalPart > 0; i += 1 {
		fractionalPart *= float64(s.base)
		result += string(symbols[int(fractionalPart)])
		_, fractionalPart = math.Modf(fractionalPart)
	}

	if isNegative {
		result = "-" + result
	}

	return result, nil
}
