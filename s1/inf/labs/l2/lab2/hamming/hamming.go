package hamming

import (
	"fmt"
)

type HammingEncoder struct {
	informationalBitsNumber int
	checkBitsNumber         int
}

func NewHammingEncoder(informationalBitsNumber int) *HammingEncoder {
	return &HammingEncoder{
		informationalBitsNumber: informationalBitsNumber,
		checkBitsNumber:         getCheckBitsNumber(informationalBitsNumber),
	}
}

type DecodingResult struct {
	Valid       bool
	Syndrome    int
	ValidString string
	Value       string
}

func (h *HammingEncoder) Decode(str string) (DecodingResult, error) {
	if len(str) != h.informationalBitsNumber+h.checkBitsNumber {
		return DecodingResult{}, fmt.Errorf("invalid string length")
	}

	s := 0

	for i := 0; i < h.checkBitsNumber; i++ {
		v := 1 << i
		m := 0
		for j := v - 1; j < len(str); j += 2 * v {
			for l := j; l <= j+v-1 && l < len(str); l += 1 {
				if d := str[l] - '0'; d == 0 || d == 1 {
					m = m ^ int(d)
				} else {
					return DecodingResult{}, fmt.Errorf("unexpected symbol: %q", str[l])
				}
			}
		}

		s += m << i
	}

	validString := []byte(str)
	if s > 0 {
		if validString[s-1] == '1' {
			validString[s-1] = '0'
		} else {
			validString[s-1] = '1'
		}
	}

	value := make([]byte, 0)

	for i := range validString {
		if (i+1)&i != 0 {
			value = append(value, validString[i])
		}
	}

	return DecodingResult{
		Valid:       s == 0,
		Syndrome:    s,
		ValidString: string(validString),
		Value:       string(value),
	}, nil
}

func getCheckBitsNumber(k int) int {
	var n = 1
	for ; (1 << n) < n+k+1; n += 1 {
	}
	return n
}
