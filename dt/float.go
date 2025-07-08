package dt

import "math"

var FloatEps = 1e-8 // tolerance

func FloatEqual(a, b float64) bool {
	return FloatEqualC(a, b, FloatEps)
}

func FloatEqualC(a, b, eps float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}

func FloatIsInt(f float64) bool {
	return FloatIsIntC(f, FloatEps)
}

func FloatIsInt64(f float64) bool {
	if f > math.MaxInt64 || f < math.MinInt64 {
		return false
	}

	return FloatIsInt(f)
}

func FloatIsUInt64(f float64) bool {
	if f > math.MaxUint64 || f < 0 {
		return false
	}

	return FloatIsInt(f)
}

func FloatToInt64(f float64) int64 {
	if f > math.MaxInt64 {
		return math.MaxInt64
	}
	if f < math.MinInt64 {
		return math.MinInt64
	}

	return int64(f)
}

func FloatToUInt64(f float64) uint64 {
	if f > math.MaxUint64 {
		return math.MaxUint64
	}
	if f < 0 {
		return 0
	}

	return uint64(f)
}

func FloatIsIntC(f, eps float64) bool {
	return FloatEqualC(math.Mod(f, 1.0), 0, eps)
}
