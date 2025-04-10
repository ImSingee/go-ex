package copier

import (
	"fmt"
	. "reflect"
	"testing"
)

func TestAnything(t *testing.T) {
	tests := []any{
		`"Now cut that out!"`,
		39,
		true,
		false,
		2.14,
		[]string{
			"Phil Harris",
			"Rochester van Jones",
			"Mary Livingstone",
			"Dennis Day",
		},
		[2]string{
			"Jell-O",
			"Grape-Nuts",
		},
	}

	for _, expected := range tests {
		actual := mustDeepCopy(expected)
		fmt.Println(actual)
	}
	// Output:
	// "Now cut that out!"
	// 39
	// true
	// false
	// 2.14
	// [Phil Harris Rochester van Jones Mary Livingstone Dennis Day]
	// [Jell-O Grape-Nuts]
}

type Foo struct {
	Foo *Foo
	Bar int
}

func TestMap(t *testing.T) {
	x := map[string]*Foo{
		"foo": &Foo{Bar: 1},
		"bar": &Foo{Bar: 2},
	}
	y := mustDeepCopy(x).(map[string]*Foo)
	for _, k := range []string{"foo", "bar"} { // to ensure consistent order
		fmt.Printf("x[\"%v\"] = y[\"%v\"]: %v\n", k, k, x[k] == y[k])
		fmt.Printf("x[\"%v\"].Foo = y[\"%v\"].Foo: %v\n", k, k, x[k].Foo == y[k].Foo)
		fmt.Printf("x[\"%v\"].Bar = y[\"%v\"].Bar: %v\n", k, k, x[k].Bar == y[k].Bar)
	}
	// Output:
	// x["foo"] = y["foo"]: false
	// x["foo"].Foo = y["foo"].Foo: true
	// x["foo"].Bar = y["foo"].Bar: true
	// x["bar"] = y["bar"]: false
	// x["bar"].Foo = y["bar"].Foo: true
	// x["bar"].Bar = y["bar"].Bar: true
}

func TestInterface(t *testing.T) {
	x := []any{nil}
	y := mustDeepCopy(x).([]any)
	if !DeepEqual(x, y) || len(y) != 1 {
		t.Errorf("expect %v == %v; y had length %v (expected 1)", x, y, len(y))
	}
	var a any
	b := mustDeepCopy(a)
	if a != b {
		t.Errorf("expected %v == %v", a, b)
	}
}

func TestAvoidInfiniteLoops(t *testing.T) {
	x := &Foo{
		Bar: 4,
	}
	x.Foo = x
	y := mustDeepCopy(x).(*Foo)
	fmt.Printf("x == y: %v\n", x == y)
	fmt.Printf("x == x.Foo: %v\n", x == x.Foo)
	fmt.Printf("y == y.Foo: %v\n", y == y.Foo)
	// Output:
	// x == y: false
	// x == x.Foo: true
	// y == y.Foo: true
}

func TestUnsupportedKind(t *testing.T) {
	x := func() {}

	tests := []any{
		x,
		map[bool]any{true: x},
		[]any{x},
	}

	for _, test := range tests {
		y, err := deepCopy(test)
		if y != nil {
			t.Errorf("expected %v to be nil", y)
		}
		if err == nil {
			t.Errorf("expected err to not be nil")
		}
	}
}

func TestUnsupportedKindPanicsOnMust(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a panic; didn't get one")
		}
	}()
	x := func() {}
	mustDeepCopy(x)
}

func TestMismatchedTypesFail(t *testing.T) {
	tests := []struct {
		input any
		kind  Kind
	}{
		{
			map[int]int{1: 2, 2: 4, 3: 8},
			Map,
		},
		{
			[]int{2, 8},
			Slice,
		},
	}
	for _, test := range tests {
		for kind, copier := range copiers {
			if kind == test.kind {
				continue
			}
			actual, err := copier(test.input, nil)
			if actual != nil {

				t.Errorf("%v attempted value %v as %v; should be nil value, got %v", test.kind, test.input, kind, actual)
			}
			if err == nil {
				t.Errorf("%v attempted value %v as %v; should have gotten an error", test.kind, test.input, kind)
			}
		}
	}
}

func TestTwoNils(t *testing.T) {
	type Foo struct {
		A int
	}
	type Bar struct {
		B int
	}
	type FooBar struct {
		Foo  *Foo
		Bar  *Bar
		Foo2 *Foo
		Bar2 *Bar
	}

	src := &FooBar{
		Foo2: &Foo{1},
		Bar2: &Bar{2},
	}

	dst := mustDeepCopy(src)

	if !DeepEqual(src, dst) {
		t.Errorf("expect %v == %v; ", src, dst)
	}

}
