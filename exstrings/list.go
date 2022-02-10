package exstrings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImSingee/go-ex/set"
)

// ListToInt 将 []string 转换为 []int, 忽略前后空格
func ListToInt(ss []string) ([]int, error) {
	result := make([]int, len(ss))

	for i, s := range ss {
		v, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return nil, fmt.Errorf("%dth element cannot convert to int: %w", i, err)
		}

		result[i] = v
	}

	return result, nil
}

// ListToInterface 将 []string 转换为 []interface
func ListToInterface(ss []string) []interface{} {
	result := make([]interface{}, len(ss))

	for i, s := range ss {
		result[i] = s
	}

	return result
}

// Diff 给出两个数组，返回二者的差异
// 返回值 1 为 cmp 比 base 增加的元素
// 返回值 2 为 cmp 比 base 减少的元素
// 这一函数的时间和空间复杂度均为 O(M+N)
func Diff(cmp, base []string) (add, sub, equal []string) {
	cmpSet := set.NewFromString(cmp)
	baseSet := set.NewFromString(base)

	add, _ = cmpSet.Difference(baseSet).ToStringList()
	sub, _ = baseSet.Difference(cmpSet).ToStringList()
	equal, _ = baseSet.Intersection(cmpSet).ToStringList()

	return
}

func Merge(list ...[]string) []string {
	n := 0
	for _, l := range list {
		n += len(l)
	}

	result := make([]string, 0, n)
	for _, l := range list {
		result = append(result, l...)
	}

	return result
}

func DeDuplicate(base []string) []string {
	if len(base) <= 1 {
		return base
	}

	newSlice := make([]string, 0, len(base))
	s := set.New()

	for _, e := range base {
		if s.Has(e) {
			continue
		}

		newSlice = append(newSlice, e)
		s.Insert(e)
	}

	return newSlice
}

func GetIntersection(sliceList ...[]string) []string {
	sets := make([]*set.Set, len(sliceList))
	for i, e := range sliceList {
		sets[i] = set.NewFromString(e)
	}

	ee, _ := set.Intersection(sets...).ToStringList()
	return ee
}
