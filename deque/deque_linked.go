package deque

import (
	"container/list"
	"errors"
)

var ErrEmptyLinkedDeque = errors.New("LinkedDeque is empty")

type LinkedDeque[T any] struct {
	container *list.List
}

func (d *LinkedDeque[T]) init() {
	if d.container == nil {
		d.container = list.New()
	}
}

func (d *LinkedDeque[T]) Size() int {
	d.init()
	return d.container.Len()
}

func (d *LinkedDeque[T]) AddFront(itm T) {
	d.init()
	d.container.PushFront(itm)
}

func (d *LinkedDeque[T]) AddTail(itm T) {
	d.init()
	d.container.PushBack(itm)
}

func (d *LinkedDeque[T]) RemoveFront() (T, error) {
	d.init()
	if d.container.Len() == 0 {
		var zero T
		return zero, ErrEmptyLinkedDeque
	}

	elem := d.container.Front()
	d.container.Remove(elem)
	return elem.Value.(T), nil
}

func (d *LinkedDeque[T]) RemoveTail() (T, error) {
	d.init()
	if d.container.Len() == 0 {
		var zero T
		return zero, ErrEmptyLinkedDeque
	}

	elem := d.container.Back()
	d.container.Remove(elem)
	return elem.Value.(T), nil
}
