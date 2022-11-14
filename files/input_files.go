package files

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const CharactersPerLine = 27

func readFile(inputFile string) (lines []string, err error) {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f.Close()
	return
}

func groupInputLines(lines []string) (configurations [][3]string) {
	for i := 0; i < len(lines); i += 4 {
		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]
		configurations = append(configurations, [3]string{line1, line2, line3})
	}
	return
}

func parseInputLinesToBuildListDecodedNumbers(configurations [][3]string) (inputDecodedNumbers [][]CodedNumber) {
	for _, account := range configurations {
		top := account[0]
		middle := account[1]
		base := account[2]

		if len(top) != CharactersPerLine ||
			len(middle) != CharactersPerLine ||
			len(base) != CharactersPerLine {
			//TODO: What I will do with this case
		}

		var acD []CodedNumber
		for i := 0; i < CharactersPerLine; i += 3 {
			cd := top[i:i+3] + middle[i:i+3] + base[i:i+3]
			decoded := identifyNumberWithCodedChars(cd)
			acD = append(acD, decoded)
		}
		inputDecodedNumbers = append(inputDecodedNumbers, acD)
	}
	return
}

func identifyNumberWithCodedChars(coded string) CodedNumber {
	for _, cd := range codedNumbers {
		if cd.Characters == coded {
			return cd
		}
	}
	return CodedNumber{
		Number:     -1,
		Characters: coded,
	}
}

func transformDecodedToAccount(inputDecodedNumbers [][]CodedNumber) (accounts []string) {
	for _, ia := range inputDecodedNumbers {
		var numberAcc = ""
		for _, v := range ia {
			if v.Number > -1 {
				numberAcc += strconv.Itoa(v.Number)
			} else {
				numberAcc += "?"
			}

		}
		accounts = append(accounts, numberAcc)
	}
	return
}

func ParseChallengeIn(inputFile string) (accounts []string, err error) {
	lines, err := readFile(inputFile)
	if err != nil {
		return accounts, err
	}
	configurations := groupInputLines(lines)
	inputDecodedNumbers := parseInputLinesToBuildListDecodedNumbers(configurations)
	return transformDecodedToAccount(inputDecodedNumbers), nil
}
