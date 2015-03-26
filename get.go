package main

import (
	"fmt"
	"strings"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/coreos/go-etcd/etcd"
	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/yasuyuky/jsonpath"
)

type getCmd struct {
	key  string
	path string
}

func newGetCmd(c *cli.Context) (*getCmd, error) {
	args := c.Args()
	if len(args) < 2 {
		return nil, fmt.Errorf("format is get KEY PATH")
	}
	return &getCmd{key: args[0], path: args[1]}, nil
}

func get(c *cli.Context) {
	cmd, err := newGetCmd(c)
	if err != nil {
		exitln(err)
	}
	cl := etcd.NewClient(c.GlobalStringSlice("servers"))
	resp, err := cl.Get(cmd.key, false, false)

	if err != nil {
		exitln(err)
	}
	data, err := jsonpath.DecodeString(resp.Node.Value)
	if err != nil {
		exitln(err)
	}

	strSlice := strings.Split(cmd.path, ".")
	pathSlice := make([]interface{}, len(strSlice))
	for i, str := range strSlice {
		pathSlice[i] = str
	}

	res, err := jsonpath.Get(data, pathSlice, nil)
	if err != nil {
		exitln("path not found")
	}
	if res == nil {
		exitf("no value at %+v", pathSlice)
	}
	loglnf("%+v", res)
}
