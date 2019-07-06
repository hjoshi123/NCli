package main

import (
	"fmt"
	"net"
	"os"

	"github.com/urfave/cli"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	slogger := logger.Sugar()
	slogger.Debug("Application started")

	app := cli.NewApp()
	app.Name = "ncli"
	app.Usage = "Lets you query IPs, CNAMEs"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "nuvonic.net",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Look up Name servers for given host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))

				if err != nil {
					slogger.Errorf("Error occured in ns %v", err)
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cname)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		slogger.Fatalf("Application exited due to %v", err)
	}
	slogger.Debug("Application Ended")
}
