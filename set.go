package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
)

type setCmd struct {
	key  string
	val  string
	path string
}

func newSetCmd(c *cli.Context) (*setCmd, error) {
	args := c.Args()
	if len(args) < 3 {
		return nil, fmt.Errorf("format is set KEY PATH VAL")
	}
	return &setCmd{key: args[0], val: args[2], path: args[1]}, nil
}

func set(c *cli.Context) {
	ttl := c.Int("ttl")
	cmd, err := newSetCmd(c)
	if err != nil {
		exitln(err)
	}
	cl := etcd.NewClient(c.GlobalStringSlice("servers"))
	resp, err := cl.Get(cmd.key, false, false)
	if err != nil {
		exitln("key doesn't exist")
	}
	body := resp.Node.Value

	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		exitln(err)
	}

	split := strings.Split(cmd.path, ".")
	newMap, err := setInMap(m, split, cmd.val)
	if err != nil {
		exitln(err)
	}

	b, err := json.Marshal(newMap)
	if err != nil {
		exitln(err)
	}
	_, err = cl.Set(cmd.key, string(b), uint64(ttl))
	if err != nil {
		exitln(err)
	}
	logln("success")
}
