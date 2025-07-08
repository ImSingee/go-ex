package dt

func ParseBoolFromString(b string) (v bool, ok bool) {
	switch b {
	case "1", "on", "ok", "enable", "true":
		return true, true
	case "0", "off", "no", "disable", "false", "":
		return false, true
	}

	if v, ok := IntFromString(b); ok {
		return ParseBoolFromNumber(v), true
	}

	return false, false
}

// 0 = false
// other = true
func ParseBoolFromNumber(number *GenericNumber) bool {
	if !number.IsInt64() {
		return true // Too large, or non-integer floating point
	}

	return number.Int64() != 0
}

func ParseBool(b interface{}) (v bool, ok bool) {
	if b == nil {
		return false, true
	}

	switch v := b.(type) {
	case bool:
		return v, true
	case string, Stringer:
		return ParseBoolFromString(ToString(v))
	}

	if v, ok := NumberFromBasicType(b); ok {
		return ParseBoolFromNumber(v), true
	}

	return false, false
}
