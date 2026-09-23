package converter_test

import (
	"converter"
	numeralsystems "converter/numeral_systems"
	"fmt"
	"testing"
)

func TestLab(t *testing.T) {
	converter := converter.New()
	converter.Register("dec", numeralsystems.NewStandardNumeralSystem(10))
	converter.Register("7", numeralsystems.NewStandardNumeralSystem(7))
	converter.Register("13", numeralsystems.NewStandardNumeralSystem(13))
	converter.Register("hex", numeralsystems.NewStandardNumeralSystem(16))
	converter.Register("oct", numeralsystems.NewStandardNumeralSystem(8))
	converter.Register("bin", numeralsystems.NewStandardNumeralSystem(2))
	converter.Register("5", numeralsystems.NewStandardNumeralSystem(5))
	converter.Register("fib", numeralsystems.NewFibonacciNumeralSystem())
	converter.Register("fact", numeralsystems.NewFactorialNumeralSystem())
	converter.Register("berg", numeralsystems.NewBergmanNumeralSystem())
	converter.Register("9C", numeralsystems.NewSymmetricalNumeralSystem(9))

	tests := []struct {
		value string
		from  string
		to    string
	}{
		{
			value: "80308",
			from:  "dec",
			to:    "5",
		},
		{
			value: "2C2C6",
			from:  "13",
			to:    "dec",
		},
		{
			value: "55345",
			from:  "7",
			to:    "13",
		},
		{
			value: "86.72",
			from:  "dec",
			to:    "bin",
		},
		{
			value: "30.13",
			from:  "hex",
			to:    "bin",
		},
		{
			value: "20.35",
			from:  "oct",
			to:    "bin",
		},
		{
			value: "0.000111",
			from:  "bin",
			to:    "hex",
		},
		{
			value: "0.111101",
			from:  "bin",
			to:    "dec",
		},
		{
			value: "B7.D0",
			from:  "hex",
			to:    "dec",
		},
		{
			value: "100",
			from:  "dec",
			to:    "fact",
		},
		{

			value: "1000",
			from:  "dec",
			to:    "berg",
		},
		{
			value: "1000101",
			from:  "berg",
			to:    "dec",
		},
		{
			value: "1000101",
			from:  "fact",
			to:    "dec",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			result, err := converter.Convert(tt.value, tt.from, tt.to)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%s in %s equals %s in %s", tt.value, converter.NumericalSystem(tt.from).Name(), result, converter.NumericalSystem(tt.to).Name())
		})
	}
}
