package main

import "errors"

var ErrInvalidPath = errors.New("invalid path")

func traverse(m interface{}, pathSplit []string) (interface{}, error) {
	if len(pathSplit) == 0 {
		return m, nil
	}
	switch t := m.(type) {
	case map[string]interface{}:
		i, ok := t[pathSplit[0]]
		if !ok {
			return nil, ErrInvalidPath
		}
		return i, nil
	default:
		return nil, ErrInvalidPath
	}
}
