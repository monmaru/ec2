package main

import (
	"fmt"
	"os"

	"github.com/monmaru/ec2/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ec2"
	app.Version = "1.0"
	app.Author = "monmaru"
	app.Description = "simple cli tool for Amazon EC2"
	app.Commands = commands
	app.CommandNotFound = usage
	app.Run(os.Args)
}

var regionFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Value: "ap-northeast-1",
		Usage: "Specify the target AWS region.",
	},
}

var commands = []cli.Command{
	{
		Name:   "list",
		Usage:  "List all EC2 instance infomation",
		Action: cmd.ListInstances,
		Flags:  regionFlags,
	},
	{
		Name:      "start",
		Usage:     "Start EC2 instances",
		Action:    cmd.StartInstances,
		ArgsUsage: "EC2 instance id listt",
		Flags:     regionFlags,
	},
	{
		Name:      "stop",
		Usage:     "Stop EC2 instances",
		Action:    cmd.StopInstances,
		ArgsUsage: "EC2 instance id listt",
		Flags:     regionFlags,
	},
}

func usage(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.\n", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
