package dt

import "reflect"

// In: Check if `expect` value is in container
//
// Note: Container's type shouyld be []T (or [N]T) and element's type must be T
func InSlice(container interface{}, expect interface{}) (exist bool) {
	containerReflectValue := reflect.ValueOf(container)

	switch containerReflectValue.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic("invalid type")
	}

	length := containerReflectValue.Len()
	for i := 0; i < length; i++ {
		e := containerReflectValue.Index(i).Interface()
		if reflect.DeepEqual(e, expect) {
			return true
		}
	}

	return false
}
