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
		cli.IntFlag{
			Name:  "ttl, t",
			Value: 0,
			Usage: "set the ttl",
		},
	}
	app.Commands = []cli.Command{
		cli.Command{
			Name:    "set",
			Aliases: []string{"s"},
			Usage:   "set a JSON value",
			Action:  set,
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
