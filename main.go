package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

var ErrNotEnoughArgs = errors.New("not enough arguments")

const Sep = "."

func keyAndVal(c *cli.Context) (string, string, error) {
	args := c.Args()
	if len(args) < 2 {
		return "", "", ErrNotEnoughArgs
	}
	return args[0], args[1], nil
}

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
			Usage:   "set a JSON value",
			Action: func(c *cli.Context) {
				key, val, err := keyAndVal(c)
				if err != nil {
					log.Fatal(err)
				}
				cl := etcd.NewClient(c.GlobalStringSlice("servers"))
				resp, err := cl.Get(key, false, false)
				if err != nil {
					log.Fatal(err)
				}
				var m = map[string]interface{}{}
				err = json.Unmarshal(resp.Body, &m)
				if err != nil {
					log.Fatal(err)
				}
				split := strings.Split(val, Sep)
				ret, err := traverse(m, split)
				//write it back now

			},
		},
		cli.Command{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get a JSON value",
			Action: func(c *cli.Context) {
				fmt.Printf("servers: %+v\n", c.String("servers"))
				fmt.Printf("get value %+v\n", c.String("get"))
			},
		},
	}
	app.Run(os.Args)

}
