package dt

import (
	"fmt"
	"strconv"
)

type Stringer = fmt.Stringer

type String string

func (s String) String() string {
	return string(s)
}

func ToString(v interface{}) string {
	switch vv := v.(type) {
	case string:
		return vv
	case Stringer: // includes *GenericNumber
		return vv.String()
	case byte: // & uint8
		return string(vv)
	case rune: // & int32
		return string(vv)
	case float64:
		return strconv.FormatFloat(vv, 'f', -1, 64)
	case int, int8, int16, int64, uint, uint16, uint32, uint64:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("%#+v", v)
	}
}

func ToStringer(v interface{}) Stringer {
	switch vv := v.(type) {
	case Stringer:
		return vv
	default:
		return String(ToString(v))
	}
}
