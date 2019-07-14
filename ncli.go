package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/hjoshi123/NCli/portscanner"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
)

func main() {
	logger, _ := zap.NewDevelopment()
	slogger := logger.Sugar()
	if os.Getenv("DEV") == "dev" {
		slogger.Debug("Application started")
	}

	app := cli.NewApp()
	app.Name = "ncli"
	app.Version = "0.0.2"
	app.Usage = "Lets you query IPs, CNAMEs"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "nuvonic.net",
		},
	}

	portFlag := []cli.Flag{
		cli.IntFlag{
			Name:  "port1",
			Value: 8080,
			Usage: "Enter the beginning port range",
		},
		cli.IntFlag{
			Name:  "port2",
			Value: 8082,
			Usage: "Enter the ending port range",
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
					if os.Getenv("DEV") == "dev" {
						slogger.Errorf("Error occured in ns %v", err)
					}
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
					if os.Getenv("DEV") == "dev" {
						slogger.Errorf("Error occured in ip %v", err)
					}
					return err
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
					if os.Getenv("DEV") == "dev" {
						slogger.Errorf("Error occured in cname %v", err)
					}
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX (Mail Exchange) records for a particular host",
			Flags: flags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					if os.Getenv("DEV") == "dev" {
						slogger.Errorf("Error occured in mx %v", err)
					}
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
		{
			Name:  "ha",
			Usage: "Looks up the given host and gives out host's addresses. ",
			Flags: flags,
			Action: func(c *cli.Context) error {
				addrs, err := net.LookupHost(c.String("host"))
				if err != nil {
					if os.Getenv("DEV") == "dev" {
						slogger.Errorf("Error occured in ha %v", err)
					}
					return err
				}
				for i := 0; i < len(addrs); i++ {
					fmt.Println(addrs[i])
				}
				return nil
			},
		},
		{
			Name:  "port",
			Usage: "Checks if the given range of ports at localhost is open or closed",
			Flags: portFlag,
			Action: func(c *cli.Context) error {
				ps := &portscanner.PortScanner{
					IP:   "127.0.0.1",
					Lock: semaphore.NewWeighted(portscanner.Ulimit()),
				}
				ps.Start(c.Int("port1"), c.Int("port2"), 500*time.Millisecond)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Application exited due to %v", err)
	}
	if os.Getenv("DEV") == "dev" {
		slogger.Debug("Application Ended")
	}
}
