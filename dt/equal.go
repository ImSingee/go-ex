package dt

import "reflect"

func Equal(expect, real interface{}) bool {
	if reflect.DeepEqual(expect, real) {
		return true
	}

	switch ev := expect.(type) {
	// TODO
	case bool:
		rv, ok := ParseBool(real)
		if !ok {
			return false
		}
		return ev == rv
	default:
		return false
	}
}
