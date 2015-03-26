package main

import (
	"fmt"
	"reflect"
)

func noKeyErr(key string) error {
	return fmt.Errorf("no key %s found", key)
}

func keyExistsIface(m *map[string]interface{}, key string) bool {
	_, ok := (*m)[key]
	return ok
}
func keyExistsInt(m *map[string]int, key string) bool {
	_, ok := (*m)[key]
	return ok
}
func keyExistsStr(m *map[string]string, key string) bool {
	_, ok := (*m)[key]
	return ok
}

// getPtr gets a pointer to the value in the map m at the given path.
// if the path doesn't exist in m, returns nil and an error indicating so.
// if no error, getPtr will return a pointer to the value at path.
// you can follow the pointer and set it to the value you want. you
// should type check it first
func setInMap(m interface{}, path []string, newVal interface{}) (interface{}, error) {
	// terminal types
	if len(path) == 0 {
		return newVal, nil
	}

	key := path[0]
	termErr := fmt.Errorf("path terminates early at %s", key)
	rem := path[1:]

	val := reflect.ValueOf(m)
	switch val.Kind() {
	case reflect.Map:
		keys := val.MapKeys()
		found := false
		for _, k := range keys {
			if reflect.DeepEqual(k.Interface(), key) {
				found = true
				break
			}
		}
		if !found {
			return nil, termErr
		}
		subMap := val.MapIndex(reflect.ValueOf(key))
		newSubMap, err := setInMap(subMap.Interface(), rem, newVal)
		if err != nil {
			return nil, err
		}
		val.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(newSubMap))
		return val.Interface(), nil
	default:
		return nil, fmt.Errorf("value at path %s is unsupported type %T", key, m)
	}
}
