package accounts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_verificateAccount(t *testing.T) {
	input := map[string]bool{
		"000000000": true,
		"111111111": false,
		"222222222": false,
		"333333333": false,
		"444444444": false,
		"555555555": false,
		"666666666": false,
		"345882865": true,
		"777777777": false,
		"888888888": false,
		"999999999": false,
		"123456789": true, // 9 + 16 + 21 + 24 + 25 + 24 + 21 + 16 + 9  = 165/11 = 15, mod 0
		"000000051": true, // 1 + 10 = 11 / 11 = 1, mod 0
	}

	for k, v := range input {
		r := verificateAccount(k)
		assert.Equal(t, v, r, "The value is giving problems ", k)
	}

}

func Test_ValidateInputData(t *testing.T) {
	expected := map[string][]bool{
		"000000000": []bool{true, true},
		"111111111": []bool{true, false},
		"222222222": []bool{true, false},
		"333333333": []bool{true, false},
		"444444444": []bool{true, false},
		"555555555": []bool{true, false},
		"666666666": []bool{true, false},
		"777777777": []bool{true, false},
		"888888888": []bool{true, false},
		"999999999": []bool{true, false},
		"123456789": []bool{true, true}, // 9 + 16 + 21 + 24 + 25 + 24 + 21 + 16 + 9  = 165/11 = 15, mod 0
		"000000051": []bool{true, true}, // 1 + 10 = 11 / 11 = 1, mod 0
		"345882865": []bool{true, true},
		"49006771?": []bool{false, false},
		"1234?678?": []bool{false, false},
	}

	keys := make([]string, len(expected))
	i := 0
	for k := range expected {
		keys[i] = k
		i++
	}

	va := ValidateInputData(keys)
	for _, ra := range va {
		expR := expected[ra.InputNumber]
		assert.Equal(t, expR[0], ra.ValidNumber, "The Account is failing on the Validation for numbers")
		assert.Equal(t, expR[1], ra.VerifiedAcc, "The Account is failing on the Verification")
	}

}
