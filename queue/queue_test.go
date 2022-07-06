package queue

import "testing"
import "container/list"
import "math/rand"

func ensureEmpty(t *testing.T, q *Queue[int]) {
	if l := q.Len(); l != 0 {
		t.Errorf("q.Len() = %d, want %d", l, 0)
	}
	if e := q.Front(); e != 0 {
		t.Errorf("q.Front() = %v, want %v", e, nil)
	}
	if e := q.Back(); e != 0 {
		t.Errorf("q.Back() = %v, want %v", e, nil)
	}
	if e := q.PopFront(); e != 0 {
		t.Errorf("q.PopFront() = %v, want %v", e, nil)
	}
	if e := q.PopBack(); e != 0 {
		t.Errorf("q.PopBack() = %v, want %v", e, nil)
	}
}

func TestNew(t *testing.T) {
	q := New[int](0)
	ensureEmpty(t, q)
}

func ensureSingleton(t *testing.T, q *Queue[int]) {
	if l := q.Len(); l != 1 {
		t.Errorf("q.Len() = %d, want %d", l, 1)
	}
	if e := q.Front(); e != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
	if e := q.Back(); e != 42 {
		t.Errorf("q.Back() = %v, want %v", e, 42)
	}
}

func TestSingleton(t *testing.T) {
	q := New[int](0)
	ensureEmpty(t, q)
	q.PushFront(42)
	ensureSingleton(t, q)
	q.PopFront()
	ensureEmpty(t, q)
	q.PushBack(42)
	ensureSingleton(t, q)
	q.PopBack()
	ensureEmpty(t, q)
	q.PushFront(42)
	ensureSingleton(t, q)
	q.PopBack()
	ensureEmpty(t, q)
	q.PushBack(42)
	ensureSingleton(t, q)
	q.PopFront()
	ensureEmpty(t, q)
}

func TestDuos(t *testing.T) {
	q := New[int](0)
	ensureEmpty(t, q)
	q.PushFront(42)
	ensureSingleton(t, q)
	q.PushBack(43)
	if l := q.Len(); l != 2 {
		t.Errorf("q.Len() = %d, want %d", l, 2)
	}
	if e := q.Front(); e != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
	if e := q.Back(); e != 43 {
		t.Errorf("q.Back() = %v, want %v", e, 43)
	}
}

func ensureLength(t *testing.T, q *Queue[int], len int) {
	if l := q.Len(); l != len {
		t.Errorf("q.Len() = %d, want %d", l, len)
	}
}
func TestGrowShrink1(t *testing.T) {
	q := New[int](0)
	for i := 0; i < size; i++ {
		q.PushBack(i)
		ensureLength(t, q, i+1)
	}
	for i := 0; q.Len() > 0; i++ {
		x := q.PopFront()
		if x != i {
			t.Errorf("q.PopFront() = %d, want %d", x, i)
		}
		ensureLength(t, q, size-i-1)
	}
}
func TestGrowShrink2(t *testing.T) {
	q := New[int](0)
	for i := 0; i < size; i++ {
		q.PushFront(i)
		ensureLength(t, q, i+1)
	}
	for i := 0; q.Len() > 0; i++ {
		x := q.PopBack()
		if x != i {
			t.Errorf("q.PopBack() = %d, want %d", x, i)
		}
		ensureLength(t, q, size-i-1)
	}
}

const size = 1024

func BenchmarkPushFrontQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := New[int](0)
		for n := 0; n < size; n++ {
			q.PushFront(n)
		}
	}
}
func BenchmarkPushFrontList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushFront(n)
		}
	}
}

func BenchmarkPushBackQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := New[int](0)
		for n := 0; n < size; n++ {
			q.PushBack(n)
		}
	}
}
func BenchmarkPushBackList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushBack(n)
		}
	}
}
func BenchmarkPushBackChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make(chan interface{}, size)
		for n := 0; n < size; n++ {
			q <- n
		}
		close(q)
	}
}

var rands []float32

func makeRands() {
	if rands != nil {
		return
	}
	rand.Seed(64738)
	for i := 0; i < 4*size; i++ {
		rands = append(rands, rand.Float32())
	}
}
func BenchmarkRandomQueue(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := New[int](0)
		for n := 0; n < 4*size; n += 4 {
			if rands[n] < 0.8 {
				q.PushBack(n)
			}
			if rands[n+1] < 0.8 {
				q.PushFront(n)
			}
			if rands[n+2] < 0.5 {
				q.PopFront()
			}
			if rands[n+3] < 0.5 {
				q.PopBack()
			}
		}
	}
}
func BenchmarkRandomList(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < 4*size; n += 4 {
			if rands[n] < 0.8 {
				q.PushBack(n)
			}
			if rands[n+1] < 0.8 {
				q.PushFront(n)
			}
			if rands[n+2] < 0.5 {
				if e := q.Front(); e != nil {
					q.Remove(e)
				}
			}
			if rands[n+3] < 0.5 {
				if e := q.Back(); e != nil {
					q.Remove(e)
				}
			}
		}
	}
}

func BenchmarkGrowShrinkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := New[int](0)
		for n := 0; n < size; n++ {
			q.PushBack(i)
		}
		for n := 0; n < size; n++ {
			q.PopFront()
		}
	}
}
func BenchmarkGrowShrinkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushBack(i)
		}
		for n := 0; n < size; n++ {
			if e := q.Front(); e != nil {
				q.Remove(e)
			}
		}
	}
}
