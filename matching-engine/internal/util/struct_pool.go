package util

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func NewPool[T any]() {

}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}

func (p *Pool[T]) Put(value T) {
	p.p.Put(value)
}
