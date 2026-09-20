package numeralsystems

import (
	"fmt"
	"math"
	"strings"
)

type SymmetricalNumeralSystem struct {
	base      int
	precision int
}

func NewSymmetricalNumeralSystem(base int) *SymmetricalNumeralSystem {
	return &SymmetricalNumeralSystem{
		base:      base,
		precision: 8,
	}
}

func (s *SymmetricalNumeralSystem) Name() string {
	return fmt.Sprintf("Symmetrical{%d}", s.base)
}

func (s *SymmetricalNumeralSystem) Base() any {
	return s.base
}

func (s *SymmetricalNumeralSystem) Parse(str string) (float64, error) {
	if s.base%2 == 0 || s.base < 3 || s.base > 71 {
		return 0, fmt.Errorf("invalid base: %d", s.base)
	}

	str = strings.ToLower(str)

	values := make([]int, 0)
	dotIndex := -1
	wasMinus := false
	for _, ch := range str {
		if ch == '.' {
			if wasMinus {
				return 0, fmt.Errorf("unexpected symbol: %q", ch)
			}

			if dotIndex == -1 {
				dotIndex = len(values)
			} else {
				return 0, fmt.Errorf("unexpected symbol: %q", ch)
			}
		}

		if ch == '-' {
			if wasMinus {
				return 0, fmt.Errorf("unexpected symbol: %q", ch)
			} else {
				wasMinus = true
			}
		}

		if v, ok := charValues[ch]; ok {
			if v > s.base/2 {
				return 0, fmt.Errorf("invalid digit: %q", ch)
			}

			if wasMinus {
				values = append(values, -v)
				wasMinus = false
			} else {
				values = append(values, v)
			}
		}
	}

	if dotIndex == -1 {
		dotIndex = len(values)
	}

	var result float64

	for i, v := range values {
		result += float64(v) * math.Pow(float64(s.base), float64(dotIndex-i-1))
	}

	return result, nil
}

func (s *SymmetricalNumeralSystem) Format(n float64) (string, error) {
	if s.base%2 == 0 || s.base < 3 || s.base > 71 {
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

	for integerPart != 0 {
		v := integerPart % s.base
		integerPart /= s.base
		if v > s.base/2 {
			v -= s.base
			integerPart += 1
		}

		result = string(symbols[int(math.Abs(float64(v)))]) + result

		if v < 0 {
			result = "-" + result
		}
	}

	_, fractionalPart := math.Modf(n)
	if fractionalPart != 0 {
		return "", fmt.Errorf("fractional part currently not supported for symmetrical systems")
	}

	if isNegative {
		negativeResult := ""
		wasMinus := false
		for _, ch := range result {
			if ch == '-' {
				wasMinus = true
			} else {
				if ch == '0' || wasMinus {
					negativeResult += string(ch)
				} else {
					negativeResult += "-" + string(ch)
				}
				wasMinus = false
			}
		}

		return negativeResult, nil
	}

	return result, nil
}
