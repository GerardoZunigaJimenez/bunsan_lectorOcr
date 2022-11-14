package main

import (
	"bunsan_lectorOcr/cmd"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Lector OCR"
	app.Usage = "Utility to read Account Banks"

	app.Commands = []cli.Command{
		cmd.NewOcrLector(),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
