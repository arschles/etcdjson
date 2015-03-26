package main

import (
	"testing"

	"github.com/arschles/etcdjson/Godeps/_workspace/src/github.com/arschles/assert"
)

func TestGetPtrStr(t *testing.T) {
	origStr := "b"
	newStr := "c"
	m := map[string]map[string]string{
		"a": map[string]string{
			"b": "b",
		},
	}
	var mInter interface{} = m
	ptr, err := getPtr(mInter, []string{"a", "b"})
	assert.NoErr(t, err)
	strPtr, ok := ptr.(*string)
	assert.True(t, ok, "returned pointer was not a *string")
	assert.Equal(t, *strPtr, origStr, "returned *interface")
	*strPtr = newStr
	assert.Equal(t, strPtr, &newStr, "address of string pointer")
	assert.Equal(t, *strPtr, newStr, "actual value of string at address")

	assert.Equal(t, m["a"]["b"], newStr, "value at a.b")
}
