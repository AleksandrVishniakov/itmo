package numeralsystems

import (
	"converter/sequences"
	"fmt"
	"strconv"
)

type FactorialNumeralSystem struct {
	fact *sequences.Factorial
}

func NewFactorialNumeralSystem() *FactorialNumeralSystem {
	return &FactorialNumeralSystem{
		fact: sequences.NewFactorial(),
	}
}

func (f *FactorialNumeralSystem) Base() any {
	return nil
}

func (f *FactorialNumeralSystem) Name() string {
	return "Factorial"
}

func (f *FactorialNumeralSystem) Parse(s string) (float64, error) {
	var result int

	for i := 1; i <= len(s); i += 1 {
		if s[len(s)-i] >= '0' && s[len(s)-1] <= '9' {
			result += int(s[len(s)-i]-'0') * f.fact.Get(i)
		} else {
			return 0, fmt.Errorf("unexpected symbol: %q", s[len(s)-i])
		}
	}

	return float64(result), nil
}

func (f *FactorialNumeralSystem) Format(n float64) (string, error) {
	var result string

	v := int(n)

	for i := 2; v > 0; i++ {
		result = strconv.Itoa(v%i) + result
		v /= i
	}

	return result, nil
}
