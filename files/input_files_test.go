package files

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strconv"
	"testing"
)

func Test_readfile(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines, err := readFile(path + "/testFiles/entradas.txt")
	assert.NotEmpty(t, lines, "The file should to have at least 1 line")
	assert.NoError(t, err, "There are not expected errors on this unit case")
	for _, l := range lines {
		fmt.Println(l)
	}
}

func Test_groupInputLines(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines, err := readFile(path + "/testFiles/entradas.txt")
	configurations := groupInputLines(lines)
	assert.NotEmpty(t, configurations, "At least one configuration")
	for _, c := range configurations {
		assert.Len(t, c, 3, "Each Configuration should to have 3 lines at least")
		fmt.Println(c)
	}
}

func Test_parseInputLinesToBuildListDecodedNumbers(t *testing.T) {
	expAcc := []string{
		"000000000",
		"111111111",
		"222222222",
		"333333333",
		"444444444",
		"555555555",
		"666666666",
		"777777777",
		"888888888",
		"999999999",
		"123456789",
		"000000051",
		"49006771-1",
		"123-1-1678-1",
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines, err := readFile(path + "/testFiles/entradas.txt")
	configurations := groupInputLines(lines)
	inputDecodedNumbers := parseInputLinesToBuildListDecodedNumbers(configurations)

	assert.NotEmpty(t, inputDecodedNumbers)
	for i, ia := range inputDecodedNumbers {
		assert.NotEmpty(t, ia)
		var numberAcc = ""
		for _, v := range ia {
			numberAcc += strconv.Itoa(v.Number)
		}
		assert.Equal(t, expAcc[i], numberAcc)
	}
}

func Test_transformDecodedToAccount(t *testing.T) {
	expAcc := []string{
		"000000000",
		"111111111",
		"222222222",
		"333333333",
		"444444444",
		"555555555",
		"666666666",
		"777777777",
		"888888888",
		"999999999",
		"123456789",
		"000000051",
		"49006771?",
		"123??678?",
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	lines, err := readFile(path + "/testFiles/entradas.txt")
	configurations := groupInputLines(lines)
	inputDecodedNumbers := parseInputLinesToBuildListDecodedNumbers(configurations)
	assert.NotEmpty(t, inputDecodedNumbers)
	accounts := transformDecodedToAccount(inputDecodedNumbers)

	for i, acc := range accounts {
		assert.Equal(t, expAcc[i], acc)
	}

}
