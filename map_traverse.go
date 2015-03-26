package main

import "fmt"

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
	fmt.Printf("getPtr m = %+v (%T)\n", m, m)
	fmt.Printf("getPtr path = %+v\n", path)

	// terminal types
	if len(path) == 0 {
		fmt.Printf("terminal type %+v (%T)", m, m)
		switch t := m.(type) {
		case string:
			return &t, nil
		case float64:
			return &t, nil
		case bool:
			return &t, nil
		case map[string]interface{}:
			return &t, nil
		default:
			return nil, fmt.Errorf("unknown terminal type %+v (%T)", t, t)
		}
	}

	// intermediate types
	switch t := m.(type) {
	case map[string]interface{}:
		fmt.Println("got map[string]interface{}")
		if keyExistsIface(t, path[0]) {
			return getPtr(t[path[0]], path[1:])
		}
		return nil, noKeyErr(path[0])
	case map[string]int:
		fmt.Println("got map[string]int")
		if keyExistsInt(t, path[0]) {
			return getPtr(t[path[0]], path[1:])
		}
		return nil, noKeyErr(path[0])
	case map[string]string:
		fmt.Println("got map[string]string")
		if keyExistsStr(t, path[0]) {
			return getPtr(t[path[0]], path[1:])
		}
		return nil, noKeyErr(path[0])
	default:
		fmt.Println("got nothing")
		return nil, fmt.Errorf("path %+v doesn't exist", path)
	}
}
