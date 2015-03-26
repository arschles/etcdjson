package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/coreos/go-etcd/etcd"
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

func noKeyErr(key string) error {
	return fmt.Errorf("no key %s found", key)
}

func keyExistsIface(m map[string]interface{}, key string) bool {
	_, ok := m[key]
	return ok
}
func keyExistsInt(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}
func keyExistsStr(m map[string]string, key string) bool {
	_, ok := m[key]
	return ok
}

// getPtr gets a pointer to the value in the map m at the given path.
// if the path doesn't exist in m, returns nil and an error indicating so.
func getPtr(m interface{}, path []string) (interface{}, error) {
	switch t := m.(type) {
	case map[string]interface{}:
		if len(path) == 1 && keyExistsIface(t, path[0]) {
			val := t[path[0]]
			return &val, nil
		} else if !keyExistsIface(t, path[0]) {
			return nil, noKeyErr(path[0])
		} else {
			return getPtr(t[path[0]], path[1:])
		}
	case map[string]int:
		if len(path) == 1 && keyExistsInt(t, path[0]) {
			val := t[path[0]]
			return &val, nil
		} else if !keyExistsInt(t, path[0]) {
			return nil, noKeyErr(path[0])
		} else {
			return getPtr(t[path[0]], path[1:])
		}
	case map[string]string:
		if len(path) == 1 && keyExistsStr(t, path[0]) {
			val := t[path[0]]
			return &val, nil
		} else if !keyExistsStr(t, path[0]) {
			return nil, noKeyErr(path[0])
		} else {
			return getPtr(t[path[0]], path[1:])
		}
	default:
		return nil, fmt.Errorf("path %+v not found", path)
	}
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
	eltPtr, err := getPtr(m, split)
	if err != nil {
		exitln(err)
	}
	switch t := eltPtr.(type) {
	case *string:
		*t = cmd.val
	case *int:
		i, err := strconv.Atoi(cmd.val)
		if err != nil {
			exitf("int found at that path and couldn't convert %s", cmd.val)
		}
		*t = i
	default:
		exitf("unsupported json at the given path [%+v] (%T)", t, t)
	}

	b, err := json.Marshal(m)
	if err != nil {
		exitln(err)
	}
	_, err = cl.Set(cmd.key, string(b), uint64(ttl))
	if err != nil {
		exitln(err)
	}
	logln("success")
}
