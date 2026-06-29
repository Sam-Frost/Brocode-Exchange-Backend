package util

import (
	"log"
	"runtime"
	"sync/atomic"
)

type RingBuffer[T any] struct {
	_    [64]byte
	head atomic.Uint64
	_    [54]byte
	tail atomic.Uint64
	_    [54]byte
	mask uint64
	buf  []T
}

// Size should be power of 2
func NewRingBuffer[T any](size uint64) *RingBuffer[T] {

	if size&(size-1) != 0 {
		log.Fatal("Invalid size of ring buffer!")
	}

	return &RingBuffer[T]{
		mask: size - 1,
		buf:  make([]T, size),
	}
}

func (rb *RingBuffer[T]) Push(item T) bool {
	head := rb.head.Load()
	tail := rb.tail.Load()

	// Ring buffer is full
	if head == tail {
		return false
	}

	rb.buf[head&rb.mask] = item
	rb.head.Add(head + 1)

	return true
}

func (rb *RingBuffer[T]) PushWait(item T) {
	for !rb.Push(item) {
		runtime.Gosched()
	}
}

func (rb *RingBuffer[T]) Pop() (T, bool) {
	head := rb.head.Load()
	tail := rb.tail.Load()

	var emptyValue T

	if head == tail {
		return emptyValue, false
	}

	item := rb.buf[tail&rb.mask]
	rb.buf[tail&rb.mask] = emptyValue
	rb.tail.Store(tail + 1)

	return item, true
}

func (rb *RingBuffer[T]) PopWait() T {
	for {
		if item, ok := rb.Pop(); ok {
			return item
		}
		runtime.Gosched()
	}
}
