package main

import (
	"os"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "etcdjson"
	app.Usage = "edit JSON in Etcd"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "servers, s",
			Value: &cli.StringSlice{},
			Usage: "set the target Etcd server(s)",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "set",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "ttl",
					Value: 1000,
					Usage: "the TTL of the key to set",
				},
			},
			Usage:  "set a JSON value",
			Action: set,
		},
		cli.Command{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get a JSON value",
			Action:  get,
		},
	}
	app.Run(os.Args)

}
