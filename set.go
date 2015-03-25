package main

import (
	"fmt"
	"strings"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
)

type setCmd struct {
	key  string
	val  string
	path string
}

func newSetCmd(c *cli.Context) (*setCmd, error) {
	args := c.Args()
	if len(args) < 3 {
		return nil, fmt.Errorf("format is set KEY VAL PATH")
	}
	return &setCmd{key: args[0], val: args[1], path: args[2]}, nil
}

func set(c *cli.Context) {
	// ttl := c.GlobalInt("ttl")
	cmd, err := newSetCmd(c)
	if err != nil {
		exitln(err)
	}
	// cl := etcd.NewClient(c.GlobalStringSlice("servers"))
	// resp, err := cl.Get(cmd.key, false, false)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := jsonpath.DecodeString(resp.Node.Value)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	split := strings.Split(cmd.path, ".")
	loglnf("%s", split)
	// ret, err := traverse(m, cmd.val, split)
	// b, err := json.Marshal(ret)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// cl.Set(cmd.key, string(b), uint64(ttl))
}
