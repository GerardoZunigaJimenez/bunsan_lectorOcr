package files

import (
	"bunsan_lectorOcr/accounts"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

func Test_createOutputStringForOutputFile(t *testing.T) {
	inputData := []accounts.ReadAccount{
		{"000000000", true, true},
		{"111111111", true, false},
		{"222222222", true, false},
		{"333333333", true, false},
		{"444444444", true, false},
		{"555555555", true, false},
		{"666666666", true, false},
		{"777777777", true, false},
		{"888888888", true, false},
		{"999999999", true, false},
		{"123456789", true, true}, // 9 + 16 + 21 + 24 + 25 + 24 + 21 + 16 + 9  = 165/11 = 15, mod 0
		{"000000051", true, true}, // 1 + 10 = 11 / 11 = 1, mod 0
		{"345882865", true, true},
		{"49006771?", false, false},
		{"1234?678?", false, false},
	}

	expectedLines := ""
	for _, v := range inputData {
		if !v.ValidNumber {
			expectedLines += fmt.Sprintln(v.InputNumber, "ILL")
		}
		if !v.VerifiedAcc {
			expectedLines += fmt.Sprintln(v.InputNumber, "ERR")
		}
		expectedLines += fmt.Sprintln(v.InputNumber, "OK")
	}
	assert.Equal(t, expectedLines, createOutputStringForOutputFile(inputData))
}

func Test_CreateChallengeOut(t *testing.T) {
	//	path, err := os.Getwd()
	//	if err != nil {
	//		log.Println(err)
	//	}

	inputData := []accounts.ReadAccount{
		{"000000000", true, true},
		{"111111111", true, false},
		{"222222222", true, false},
		{"333333333", true, false},
		{"444444444", true, false},
		{"555555555", true, false},
		{"666666666", true, false},
		{"777777777", true, false},
		{"888888888", true, false},
		{"999999999", true, false},
		{"123456789", true, true}, // 9 + 16 + 21 + 24 + 25 + 24 + 21 + 16 + 9  = 165/11 = 15, mod 0
		{"000000051", true, true}, // 1 + 10 = 11 / 11 = 1, mod 0
		{"345882865", true, true},
		{"49006771?", false, false},
		{"1234?678?", false, false},
	}

	outputFile, err := CreateChallengeOut("./testFiles/", inputData)
	assert.NoError(t, err)
	assert.True(t, deepCompare(outputFile, "./testFiles/expected.out"))

}

const chunkSize = 64000

func deepCompare(file1, file2 string) bool {
	// Check file size ...

	f1, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for {
		b1 := make([]byte, chunkSize)
		_, err1 := f1.Read(b1)

		b2 := make([]byte, chunkSize)
		_, err2 := f2.Read(b2)

		if err1 != nil || err2 != nil {
			if err1 == io.EOF && err2 == io.EOF {
				return true
			} else if err1 == io.EOF || err2 == io.EOF {
				return false
			} else {
				log.Fatal(err1, err2)
			}
		}

		if !bytes.Equal(b1, b2) {
			return false
		}
	}
}
