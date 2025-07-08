package dt

type Value interface{}
type Type uint8

const (
	InvalidType Type = iota
	NilType
	BoolType
	NumberType
	StringType
)

func AsType(value interface{}, checkType Type) (Value, bool) {
	switch checkType {
	case InvalidType:
		return nil, false
	case NilType:
		return nil, true
	case BoolType:
		b, ok := value.(bool)
		return b, ok
	case NumberType:
		return NumberFromBasicIntType(value)
	case StringType:
		s, ok := value.(string)
		return s, ok
	default:
		return nil, false
	}
}
