package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/urfave/cli" // imports as package "cli"
)

func main() {
	var (
		certificate,
		object,
		root,
		proof,
		fobject,
		fproof string
	)

	app := cli.NewApp()
	app.Name = "Verifyhash"
	app.Version = "1.0"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2016 Acronis International GmbH"
	app.Usage = "Acronis Notary verify hash CLI utility"
	app.UsageText = `verifyhash [global options]

	 Required flags: -c|-o|-fo, -p|-fp, -r`
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "cert, c",
			Value:       "",
			Usage:       "ID of certificate",
			Destination: &certificate,
		},
		cli.StringFlag{
			Name:        "object, o",
			Value:       "",
			Usage:       "\"Object\" from certificate",
			Destination: &object,
		},
		cli.StringFlag{
			Name:        "fobject, fo",
			Value:       "",
			Usage:       "Path to the file with the \"object\" from certificate",
			Destination: &fobject,
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
			Name:        "fproof, fp",
			Value:       "",
			Usage:       "Path to the file with merkle proof",
			Destination: &fproof,
		},
	}

	app.Action = func(c *cli.Context) (err error) {
		var val, b []byte
		var data dataVerification

		if len(object) != 0 && len(fobject) == 0 {
			if certificate, err = hashObject([]byte(object)); err != nil {
				color.Yellow("\n%v", err)
				return nil
			}
		}
		if len(fobject) != 0 {
			b, err = ioutil.ReadFile(fobject)
			if err != nil {
				color.Yellow("\n%v", err)
				return nil
			}
			if certificate, err = hashObject(b); err != nil {
				color.Yellow("\n%v", err)
				return nil
			}
		}
		if len(fproof) != 0 {
			b, err = ioutil.ReadFile(fproof)
			if err != nil {
				color.Yellow("\n%v", err)
				return nil
			}
			proof = string(b)
		}

		if err = data.setData(certificate, root, proof); err != nil {
			color.Yellow("\n%v", err)
			return nil
		}

		if val, err = data.getValFromTree(); err != nil || len(val) == 0 {
			color.Red("\n%v", err)
			return nil
		}

		color.Green("\nVerification successful\nValue: %s\n", string(val))

		return nil
	}

	app.Run(os.Args)
}
