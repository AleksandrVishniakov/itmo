package hamming_test

import (
	"fmt"
	"lab2/hamming"
	"reflect"
	"testing"
)

func TestHamming7_4(t *testing.T) {
	h := hamming.NewHammingEncoder(4)

	tests := []struct {
		in  string
		out hamming.DecodingResult
	}{
		{
			in:  "0110010",
			out: hamming.DecodingResult{false, 7, "0110011", "1011"},
		},
		{
			in:  "1111011",
			out: hamming.DecodingResult{false, 5, "1111111", "1111"},
		},
		{
			in:  "1000101",
			out: hamming.DecodingResult{false, 3, "1010101", "1101"},
		},
		{
			in:  "0110100",
			out: hamming.DecodingResult{false, 4, "0111100", "1100"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			ans, err := h.Decode(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(ans, tt.out) {
				t.Errorf("expected: %+v\ngot : %+v\n", tt.out, ans)
			}

			t.Logf("Code %s is %s", tt.in, ans.Value)
		})
	}
}

func TestHamming15_11(t *testing.T) {
	h := hamming.NewHammingEncoder(11)

	tests := []struct {
		in  string
		out hamming.DecodingResult
	}{
		{
			in:  "011000101100001",
			out: hamming.DecodingResult{false, 10, "011000101000001", "10011000001"},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i+1), func(t *testing.T) {
			ans, err := h.Decode(tt.in)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(ans, tt.out) {
				t.Errorf("expected: %+v\ngot : %+v\n", tt.out, ans)
			}

			t.Logf("Code %s is %s", tt.in, ans.Value)
		})
	}
}
