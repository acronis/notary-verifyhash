package main

import (
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli" // imports as package "cli"
)

func main() {
	var (
		certificate,
		root,
		proof,
		eTag string
	)

	app := cli.NewApp()
	app.Name = "Verifyhash"
	app.Version = "2.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2016 Acronis International GmbH"
	app.Usage = "Acronis Notary verify hash CLI utility"
	app.UsageText = `verifyhash [global options]

	Required flags: -c, -r, -p, -e`
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "cert, c",
			Value:       "",
			Usage:       "ID of certificate",
			Destination: &certificate,
		},
		cli.StringFlag{
			Name:        "root, r",
			Value:       "",
			Usage:       "Merkle root",
			Destination: &root,
		},
		cli.StringFlag{
			Name:        "proof, p",
			Value:       "",
			Usage:       "Merkle proof",
			Destination: &proof,
		},
		cli.StringFlag{
			Name:        "etag, e",
			Value:       "",
			Usage:       "File eTag",
			Destination: &eTag,
		},
	}

	app.Action = func(c *cli.Context) error {
		err := verifyProof([]byte(proof), root, certificate, eTag)
		if err != nil {
			color.Red("Error: %s", err.Error())
		} else {
			color.Green("Verification successful")
		}

		return nil
	}

	app.Run(os.Args)
}
