package list

import (
	"fmt"
	"github.com/ImSingee/go-ex/set"
	"strings"
)

type List[T comparable] struct {
	E []T
}

func New[T comparable]() *List[T] {
	return &List[T]{make([]T, 0)}
}

func WithCapacity[T comparable](capacity int) *List[T] {
	return &List[T]{make([]T, 0, capacity)}
}

func (l *List[T]) String() string {
	if l == nil {
		return "[]"
	}

	return fmt.Sprintf("[%v]", l.Join(", "))
}

func (l *List[T]) Len() int {
	if l == nil {
		return 0
	}

	return len(l.E)
}

func (l *List[T]) IsEmpty() bool {
	return l == nil || len(l.E) == 0
}

func (l *List[T]) Join(sep string) string {
	if l == nil {
		return ""
	}

	switch len(l.E) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", l.E[0])
	}

	ss := Map[T, string](l, func(k T) string {
		return fmt.Sprintf("%v", k)
	})

	n := len(sep) * (len(ss) - 1)
	for i := 0; i < len(ss); i++ {
		n += len(ss[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(ss[0])
	for _, s := range ss[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func (l *List[T]) Add(e ...T) {
	l.E = append(l.E, e...)
}

// AddIfNotExist like Add but 只添加不存在的元素
// 时间和空间复杂度 O(M+N)
func (l *List[T]) AddIfNotExist(e ...T) {
	if l.IsEmpty() {
		return
	} else if len(e) == 1 { // 优化至 O(1) 空间复杂度
		l.AddIfNotExistSingle(e[0])
	} else {
		toAdd := set.New(e...)
		baseSet := set.New(l.E...)

		toAdd.Do(func(k T) {
			if !baseSet.Contains(k) {
				l.Add(k)
			}
		})
	}
}

// AddIfNotExistSingle like AddIfNotExist but 只能添加单个元素
// 相比于 AddIfNotExist 时间复杂度 O(N) 不变, 空间复杂度优化至 O(1) （不构建 set）
func (l *List[T]) AddIfNotExistSingle(e T) {
	if !l.IsEmpty() {
		found := false
		for _, cur := range l.E {
			if cur == e {
				found = true
				break
			}
		}
		if !found {
			l.Add(e)
		}
	}
}

// DeleteOne 删除元素 e，至多删除一个，原地操作，时间复杂度 O(N)
func (l *List[T]) DeleteOne(e T) {
	if l == nil {
		return
	}

	found := -1
	for i, s := range l.E {
		if s == e {
			found = i
			break
		}
	}

	if found != -1 { //  found
		l.E = append(l.E[:found], l.E[found+1:]...)
	}
}

// Contains 判断 l 是否包含元素 e
// 时间复杂度 O(N)
func (l *List[T]) Contains(e T) bool {
	if l.Len() == 0 {
		return false
	}

	for _, s := range l.E {
		if e == s {
			return true
		}
	}

	return false
}

// ContainsAll 返回是否包含传入的所有元素
// 时间复杂度 O(N*M)，如果数据中的列表元素过多请使用 set 来替代
func (l *List[T]) ContainsAll(all []T) bool {
	if len(all) == 0 {
		return true
	}
	if l.Len() == 0 {
		return false
	}

	for _, e := range all {
		if !l.Contains(e) {
			return false
		}
	}

	return true
}

func (l *List[T]) Equal(others []T) bool {
	if len(l.E) != len(others) {
		return false
	}
	for i := range l.E {
		if l.E[i] != others[i] {
			return false
		}
	}
	return true
}

func (l *List[T]) ShallowCopy() []T {
	if l == nil || l.E == nil {
		return nil
	}

	return append([]T{}, l.E...)
}
