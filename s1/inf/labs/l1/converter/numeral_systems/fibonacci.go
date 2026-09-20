package numeralsystems

import (
	"converter/sequences"
	"fmt"
)

type FibonacciNumeralSystem struct {
	fib *sequences.Fibonacci
}

func NewFibonacciNumeralSystem() *FibonacciNumeralSystem {
	return &FibonacciNumeralSystem{
		fib: sequences.NewFibonacci(),
	}
}

func (f *FibonacciNumeralSystem) Base() any {
	return nil
}

func (f *FibonacciNumeralSystem) Name() string {
	return "Fibonacci"
}

func (f *FibonacciNumeralSystem) Parse(s string) (float64, error) {
	var result int

	for i := 1; i <= len(s); i += 1 {
		switch s[len(s)-i] {
		case '0':
			continue
		case '1':
			result += f.fib.Get(i)
		default:
			return 0, fmt.Errorf("unexpected symbol: %q", s[len(s)-1])
		}
	}

	return float64(result), nil
}

func (f *FibonacciNumeralSystem) Format(n float64) (string, error) {
	var result string

	v := int(n)

	i := 1
	for ; f.fib.Get(i) < v; i++ {
	}

	if f.fib.Get(i) > v {
		i -= 1
	}

	for ; i > 0; i-- {
		if f.fib.Get(i) <= v {
			v -= f.fib.Get(i)
			result += "1"
		} else {
			result += "0"
		}
	}

	return result, nil
}
