package main

import (
	"fmt"
	"log"
	"os"

	syslog "github.com/RackSec/srslog"
	"github.com/urfave/cli"
)

func writeMessage(message string, dest string, prot string, info string,
	formatter syslog.Formatter) {
	sysLog, err := syslog.Dial(prot, dest,
		syslog.LOG_WARNING|syslog.LOG_DAEMON, info)
	if err != nil {
		log.Fatal(err)
	}
	sysLog.SetFormatter(formatter)
	sysLog.Info(message)
	fmt.Printf("sent message: %s\n in format: %s\n to: %s://%s\n", message, info, prot, dest)
}

func main() {
	var format string
	var message string
	var dest string
	var prot string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "format, f",
			Value:       "rfc5424",
			Usage:       "Which RFC format should be used [rfc5424, rfc3164, comp]",
			Destination: &format,
		},
		cli.StringFlag{
			Name:        "message, m",
			Value:       "DEFAULT MESSAGE",
			Usage:       "Message which should be sent",
			Destination: &message,
		},
		cli.StringFlag{
			Name:        "dest, d",
			Value:       "localhost:1234",
			Usage:       "Destination Syslog endpoint [localhost:1234<]",
			Destination: &dest,
		},
		cli.StringFlag{
			Name:        "prot, p",
			Value:       "udp",
			Usage:       "Network protocol which should be used [udp or tcp]",
			Destination: &prot,
		},
	}

	app.Action = func(c *cli.Context) error {

		switch {

		case format == "rfc5424":
			writeMessage(message, dest, prot, "RFC5424", syslog.RFC5424Formatter)

		case format == "rfc3164":
			writeMessage(message, dest, prot, "RFC3164", syslog.RFC3164Formatter)

		case format == "comp":
			writeMessage(message, dest, prot, "COMPATIBILITY", syslog.DefaultFormatter)

		default:
			return cli.NewExitError("no valid format", 1)

		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
