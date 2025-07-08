package dt

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

type Number interface {
	AsNumber() *GenericNumber
}

type GenericNumber struct {
	literal string

	above64bit bool // Whether it cannot be represented using 64 bits (number type is *big.Int / *big.Float)
	float      bool // Whether it is a floating-point number (number type is float64 / *big.Float)
	unsigned   bool // Whether it is unsigned (number type is uint64)

	number interface{} // Type is int64/uint64/float64/*big.Int/*big.Float
}

func ConvertSignedIntToInt64(num interface{}) (int64, bool) {
	switch v := num.(type) {
	case int:
		return int64(v), true
	case int8:
		return int64(v), true
	case int16:
		return int64(v), true
	case int32:
		return int64(v), true
	case int64:
		return v, true
	default:
		return 0, false
	}
}

func ConvertUnsignedIntToUInt64(num interface{}) (uint64, bool) {
	switch v := num.(type) {
	case uint:
		return uint64(v), true
	case uint8:
		return uint64(v), true
	case uint16:
		return uint64(v), true
	case uint32:
		return uint64(v), true
	case uint64:
		return v, true
	default:
		return 0, false
	}
}

func ConvertFloatToFloat64(num interface{}) (float64, bool) {
	switch v := num.(type) {
	case float32:
		return float64(v), true
	case float64:
		return v, true
	default:
		return 0, false
	}
}

// Parse integer from string
func IntFromString(num string) (*GenericNumber, bool) {
	v, err := strconv.ParseInt(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     v,
		}, true
	}
	return nil, false
}

func UIntFromString(num string) (*GenericNumber, bool) {
	vv, err := strconv.ParseUint(num, 10, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, true
	}
	return nil, false
}

func BigIntFromString(num string) (*GenericNumber, bool) {
	vvv, ok := new(big.Int).SetString(num, 10)
	if ok {
		return &GenericNumber{
			literal:    num,
			above64bit: true,
			float:      false,
			number:     vvv,
		}, true
	}
	return nil, false
}

func FloatFromString(num string) (*GenericNumber, bool) {
	v, err := strconv.ParseFloat(num, 64)
	if err == nil {
		return &GenericNumber{
			literal:    num,
			float:      true,
			unsigned:   false,
			above64bit: false,
			number:     v,
		}, true
	}
	return nil, false
}

func BigFloatFromString(num string) (*GenericNumber, bool) {
	f, ok := new(big.Float).SetString(num)
	if ok {
		return &GenericNumber{
			literal:    num,
			above64bit: true,
			float:      true,
			number:     f,
		}, true
	}
	return nil, false
}

// Get the corresponding numeric value from string
func NumberFromString(num string) (*GenericNumber, bool) {
	// Check if there's a decimal point (integer or decimal)
	if !strings.Contains(num, ".") {
		// Try to parse as integer
		if v, ok := IntFromString(num); ok {
			return v, true
		}
		// Try to parse as unsigned integer
		if v, ok := UIntFromString(num); ok {
			return v, true
		}
		// Try to parse as big integer
		if v, ok := BigIntFromString(num); ok {
			return v, true
		}
	} else {
		// Try to parse as decimal
		if v, ok := FloatFromString(num); ok {
			return v, true
		}

		//// Try to parse as big decimal
		//v, err = BigFloatFromString(num)
		//if err == nil {
		//	return v, nil
		//}
	}

	// Return error (not a number)
	return nil, false
}

func NumberFromBasicInt(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertSignedIntToInt64(num); ok {
		return &GenericNumber{
			literal:    strconv.FormatInt(vv, 10),
			float:      false,
			unsigned:   false,
			above64bit: false,
			number:     vv,
		}, true
	}

	return nil, false
}

func NumberFromBasicUInt(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertUnsignedIntToUInt64(num); ok {

		return &GenericNumber{
			literal:    strconv.FormatUint(vv, 10),
			float:      false,
			unsigned:   true,
			above64bit: false,
			number:     vv,
		}, true
	}
	return nil, false
}

func NumberFromBasicFloat(num interface{}) (*GenericNumber, bool) {
	if vv, ok := ConvertFloatToFloat64(num); ok {
		return &GenericNumber{
			literal:    ToString(vv),
			float:      true,
			above64bit: false,
			number:     vv,
		}, true
	}

	return nil, false
}

func NumberFromBasicIntType(num interface{}) (*GenericNumber, bool) {
	if v, ok := NumberFromBasicInt(num); ok {
		return v, true
	}

	if v, ok := NumberFromBasicUInt(num); ok {
		return v, true
	}

	return nil, false
}

func NumberFromBasicType(num interface{}) (*GenericNumber, bool) {
	if v, ok := NumberFromBasicIntType(num); ok {
		return v, true
	}

	if v, ok := NumberFromBasicFloat(num); ok {
		return v, true
	}

	return nil, false
}

func ParseNumber(num interface{}) (*GenericNumber, bool) {
	if v, ok := num.(*GenericNumber); ok {
		return v, true
	}
	if v, ok := NumberFromBasicType(num); ok {
		return v, true
	}

	switch v := num.(type) {
	case Stringer, string:
		if vv, ok := NumberFromString(ToString(v)); ok {
			return vv, true
		}

		return nil, false
	}

	return nil, false
}

// Check if the value can be converted to int64
func (num *GenericNumber) IsInt64() bool {
	if !num.float {
		if num.above64bit {
			return num.number.(*big.Int).IsInt64()
		}

		if num.unsigned {
			return num.number.(uint64) <= math.MaxInt64
		} else {
			return true
		}
	} else {
		if num.above64bit {
			return num.number.(*big.Float).IsInt()
		} else {
			return FloatIsInt64(num.number.(float64))
		}
	}
}

// Convert value to int64
// If IsInt64 == false, the returned result is undefined
func (num *GenericNumber) Int64() int64 {
	if num.float {
		return FloatToInt64(num.number.(float64))
	}
	if num.above64bit {
		return num.number.(*big.Int).Int64()
	}
	if num.unsigned {
		return int64(num.number.(uint64))
	} else {
		return num.number.(int64)
	}
}

// Check if the value can be converted to uint64
func (num *GenericNumber) IsUInt64() bool {
	if !num.float {
		if num.above64bit {
			return num.number.(*big.Int).IsUint64()
		}

		if num.unsigned {
			return true
		} else {
			return num.number.(int64) >= 0
		}
	} else {
		if num.above64bit {
			return num.number.(*big.Float).IsInt()
		} else {
			return FloatIsUInt64(num.number.(float64))
		}
	}
}

// Convert value to uint64
// If IsUInt64 == false, the returned result is undefined
func (num *GenericNumber) UInt64() uint64 {
	if num.float {
		return FloatToUInt64(num.number.(float64))
	}
	if num.above64bit {
		return num.number.(*big.Int).Uint64()
	}
	if num.unsigned {
		return num.number.(uint64)
	} else {
		return uint64(num.number.(int64))
	}
}

// Check if the value can be converted to float64
func (num *GenericNumber) IsFloat64() bool {
	if num.float {
		return !num.above64bit
	} else {
		if num.above64bit {
			// TODO
			return false
		}
		return true
	}
}

// Convert value to float64
// If IsFloat64 == false, the returned result is undefined
func (num *GenericNumber) Float64() float64 {
	if num.float {
		if num.above64bit { // Currently not using big.Float
			return 0
		}

		return num.number.(float64)
	} else {
		if num.above64bit {
			return 0
		}
		if num.unsigned {
			return float64(num.number.(uint64))
		} else {
			return float64(num.number.(int64))
		}
	}
}

func (num *GenericNumber) String() string {
	return num.literal
}

func (num *GenericNumber) Float() bool {
	return num.float
}

func (num *GenericNumber) Unsigned() bool {
	return num.unsigned
}

func (num *GenericNumber) Above64bit() bool {
	return num.above64bit
}
