package files

import (
	"bunsan_lectorOcr/accounts"
	"errors"
	"fmt"
	"os"
)

func createOutputStringForOutputFile(readAccount []accounts.ReadAccount) (lines string) {
	for _, rd := range readAccount {
		if !rd.ValidNumber {
			lines += fmt.Sprintln(rd.InputNumber, "ILL")
			continue
		}
		if !rd.VerifiedAcc {
			lines += fmt.Sprintln(rd.InputNumber, "ERR")
			continue
		}
		lines += fmt.Sprintln(rd.InputNumber, "OK")
	}
	return lines
}

func deleteIfExists(fileLocation string) error {
	_, err := os.Stat(fileLocation)
	if !errors.Is(err, os.ErrNotExist) {
		err := os.Remove(fileLocation) // remove a single file
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateChallengeOut(path string, verifiedAccounts []accounts.ReadAccount) (fileLocation string, err error) {
	pathRef, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", errors.New("the path " + path + "included is invalid")
	}
	if !pathRef.IsDir() {
		return "", errors.New("the path " + path + "is not a directory")
	}

	fileLocation = path + "/challenge.out"
	err = deleteIfExists(fileLocation)
	if err != nil {
		return fileLocation, err
	}

	f, err := os.Create(fileLocation)

	if err != nil {
		return fileLocation, err
	}
	defer f.Close()

	_, err2 := f.WriteString(createOutputStringForOutputFile(verifiedAccounts))

	if err2 != nil {
		return fileLocation, err
	}

	return fileLocation, nil
}
