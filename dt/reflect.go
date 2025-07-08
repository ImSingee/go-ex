package dt

import (
	"math"
	"reflect"
)

func MapReflectType(p reflect.Kind) Type {
	switch p {
	case reflect.Bool:
		return BoolType
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return NumberType
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return NumberType
	case reflect.Float64, reflect.Float32:
		return NumberType
	case reflect.String:
		return StringType
	default:
		return InvalidType
	}
}

var invalidReflectValue = reflect.Value{}

var reflectTypes = map[reflect.Kind]reflect.Type{
	reflect.Int:     reflect.TypeOf(1),
	reflect.Int8:    reflect.TypeOf(int8(1)),
	reflect.Int16:   reflect.TypeOf(int16(1)),
	reflect.Int32:   reflect.TypeOf(int32(1)),
	reflect.Int64:   reflect.TypeOf(int64(1)),
	reflect.Uint:    reflect.TypeOf(uint(1)),
	reflect.Uint8:   reflect.TypeOf(uint8(1)),
	reflect.Uint16:  reflect.TypeOf(uint16(1)),
	reflect.Uint32:  reflect.TypeOf(uint32(1)),
	reflect.Uint64:  reflect.TypeOf(uint64(1)),
	reflect.Float32: reflect.TypeOf(float32(1)),
	reflect.Float64: reflect.TypeOf(float64(1)),
}

func CheckOverflowFloat(x float64, k reflect.Kind) bool {
	switch k {
	case reflect.Float32:
		if x < 0 {
			x = -x
		}
		return math.MaxFloat32 < x && x <= math.MaxFloat64
	case reflect.Float64:
		return false
	}
	panic("invalid kind")
}

func CheckOverflowInt(x int64, k reflect.Kind) bool {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bitSize := reflectTypes[k].Size() * 8
		trunc := (x << (64 - bitSize)) >> (64 - bitSize)
		return x != trunc
	}
	panic("invalid kind")
}

func CheckOverflowUInt(x uint64, k reflect.Kind) bool {
	switch k {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bitSize := reflectTypes[k].Size() * 8
		trunc := (x << (64 - bitSize)) >> (64 - bitSize)
		return x != trunc
	}
	panic("invalid kind")
}

func ConvertToReflectType(v Value, toType reflect.Kind) (value reflect.Value, ok bool) {
	defer func() {
		v := recover()
		if v != nil {
			value = invalidReflectValue
			ok = false
		}
	}()

	switch vv := v.(type) {
	case bool:
		if toType == reflect.Bool {
			return reflect.ValueOf(vv), true
		} else {
			return invalidReflectValue, false
		}
	case string:
		if toType == reflect.String {
			return reflect.ValueOf(vv), true
		} else {
			return invalidReflectValue, false
		}
	case *GenericNumber:
		switch toType {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !vv.IsInt64() {
				return invalidReflectValue, false
			}
			vvv := vv.Int64()
			if CheckOverflowInt(vvv, toType) {
				return invalidReflectValue, false
			}

			return reflect.ValueOf(vvv).Convert(reflectTypes[toType]), true
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if !vv.IsUInt64() {
				return invalidReflectValue, false
			}
			vvv := vv.UInt64()
			if CheckOverflowUInt(vvv, toType) {
				return invalidReflectValue, false
			}

			return reflect.ValueOf(vvv).Convert(reflectTypes[toType]), true
		case reflect.Float32, reflect.Float64:
			if !vv.IsFloat64() {
				return invalidReflectValue, false
			}
			vvv := vv.Float64()
			if CheckOverflowFloat(vvv, toType) {
				return invalidReflectValue, false
			}

			return reflect.ValueOf(vvv).Convert(reflectTypes[toType]), true
		default:
			return invalidReflectValue, false
		}
	default:
		return invalidReflectValue, false
	}
}
