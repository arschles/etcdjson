package main

import (
	"testing"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/arschles/assert"
)

func TestGetSimple(t *testing.T) {
	newVal := "c"
	m := map[string]map[string]string{
		"a": map[string]string{
			"b": "b",
		},
	}
	newInter, err := setInMap(m, []string{"a", "b"}, newVal)
	assert.NoErr(t, err)
	newM, ok := newInter.(map[string]map[string]string)
	assert.True(t, ok, "returned interface was not a map[string]map[string]string")
	assert.Equal(t, newM["a"]["b"], newVal, "value at a.b")
	m["a"]["b"] = newVal
	assert.Equal(t, newM, m, "entire map")
}

func TestGetSimple2(t *testing.T) {
	newVal := "c"
	m := map[string]string{"a": "b"}
	newInter, err := setInMap(m, []string{"a"}, newVal)
	assert.NoErr(t, err)
	newM, ok := newInter.(map[string]string)
	assert.True(t, ok, "returned interface was not a map[string]string")
	assert.Equal(t, newM["a"], newVal, "value at a")
	m["a"] = newVal
	assert.Equal(t, newM, m, "entire map")
}
