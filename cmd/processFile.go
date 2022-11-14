package cmd

import (
	"bunsan_lectorOcr/accounts"
	"bunsan_lectorOcr/files"
	"fmt"
	"github.com/urfave/cli"
	"path/filepath"
)

func NewOcrLector() cli.Command {
	return cli.Command{
		Name:   "read",
		Usage:  "Read Account Lines ",
		Action: configureDeviceConfigurations,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "inputFilePath",
				Usage:    "File Input Location that will be parsed.  This one is mandatory",
				Required: true,
			},
		},
	}
}

func configureDeviceConfigurations(c *cli.Context) error {
	inputFilePath := c.String("inputFilePath")

	readAcounts, err := files.ParseChallengeIn(inputFilePath)
	if err != nil {
		return err
	}
	verifiedAccounts := accounts.ValidateInputData(readAcounts)

	dir := filepath.Dir(inputFilePath)
	//parent := filepath.Base(dir)
	outputFilePath, err := files.CreateChallengeOut(dir, verifiedAccounts)
	if err != nil {
		return err
	}
	fmt.Println("The ouput path is: ", outputFilePath)
	return nil
}
