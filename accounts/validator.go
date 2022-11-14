package accounts

import (
	"strconv"
)

type ReadAccount struct {
	InputNumber string
	ValidNumber bool
	VerifiedAcc bool
}

func verificateAccount(iAcc string) bool {
	r := 0
	for i, j := len(iAcc)-1, 1; i > -1; i-- {
		// -48 that means ASCII Number from 0
		asc := (int(iAcc[i]) - 48)
		rule := asc * j
		r += rule
		j++
	}
	// 1 + 10 +
	if r%11 == 0 {
		return true
	}
	return false
}

func ValidateInputData(inputAccounts []string) (verifiedAccounts []ReadAccount) {
	for _, ia := range inputAccounts {
		var ra = ReadAccount{
			InputNumber: ia,
		}
		_, err := strconv.Atoi(ia)
		if err != nil {
			verifiedAccounts = append(verifiedAccounts, ra)
			continue
		}
		if !verificateAccount(ia) {
			ra.ValidNumber = true
			verifiedAccounts = append(verifiedAccounts, ra)
			continue
		}
		ra.ValidNumber = true
		ra.VerifiedAcc = true
		verifiedAccounts = append(verifiedAccounts, ra)
	}

	return verifiedAccounts
}
