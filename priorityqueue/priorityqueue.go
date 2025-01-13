package priorityqueue

import (
	"container/heap"
	"reflect"
	"github.com/nsce9806q/javastyle-collection/util"
)

// PriorityQueue is a priority queue data structure.
type PriorityQueue[E any] struct {
	heap   *internalHeap[E]
	equals util.Equals[E]
}

// Option is a function type that sets the PriorityQueue.
type Option[E any] func(*PriorityQueue[E])

// WithCapacity is an option that sets the initial capacity.
func WithCapacity[E any](initialCapacity int) Option[E] {
	return func(pq *PriorityQueue[E]) {
		pq.heap.items = make([]E, 0, initialCapacity)
	}
}

// WithComparator is an option that sets the custom comparator.
func WithComparator[E any](comparator util.Comparator[E]) Option[E] {
	return func(pq *PriorityQueue[E]) {
		pq.heap.comparator = comparator
	}
}

// WithEquals is an option that sets the custom equality comparison function.
func WithEquals[E any](equals util.Equals[E]) Option[E] {
	return func(pq *PriorityQueue[E]) {
		pq.equals = equals
	}
}

// New creates a new PriorityQueue with the given options.
func New[E any](opts ...Option[E]) *PriorityQueue[E] {
	pq := &PriorityQueue[E]{
		heap: &internalHeap[E]{
			items:      make([]E, 0, 11),
			comparator: util.DefaultComparator[E](),
		},
	}

	for _, opt := range opts {
		opt(pq)
	}

	heap.Init(pq.heap)
	return pq
}

// Inserts the specified element into this priority queue.
// boolean add(E e)
func (pq *PriorityQueue[E]) Add(item E) bool {
	if !pq.Offer(item) {
		panic("Queue is full")
	}
	return true
}

// Inserts the specified element into this priority queue.
// boolean offer(E e)
func (pq *PriorityQueue[E]) Offer(item E) (success bool) {
	defer func() {
		if r := recover(); r != nil {
			success = false
		}
	}()
	heap.Push(pq.heap, item)
	return true
}

// Removes all of the elements from this priority queue.
// void clear()
func (pq *PriorityQueue[E]) Clear() {
	pq.heap.items = []E{}
	heap.Init(pq.heap)
}

// Returns the comparator used to order the elements in this queue, or defaultComparator if the queue uses the natural ordering of its elements.
// Comparator<? super E> comparator()
func (pq *PriorityQueue[E]) Comparator() util.Comparator[E] {
	return pq.heap.comparator
}

// Returns true if this queue contains the specified element.
// boolean contains(Object o)
func (pq *PriorityQueue[E]) Contains(item E) bool {
	// when type is comparable
	if reflect.TypeOf(item).Comparable() {
		for _, v := range pq.heap.items {
			if reflect.ValueOf(v).Interface() == reflect.ValueOf(item).Interface() {
				return true
			}
		}
		return false
	}

	// panic if not comparable and equals function is not provided
	if pq.equals == nil {
		panic("Type is not comparable and equals function is not provided")
	}

	// use equals function
	for _, v := range pq.heap.items {
		if pq.equals(v, item) {
			return true
		}
	}
	return false
}

// Retrieves and removes the head of this queue, or returns null if this queue is empty.
// E poll()
func (pq *PriorityQueue[E]) Poll() E {
	item := heap.Pop(pq.heap)
	if item == nil {
		var zero E
		return zero
	}
	return item.(E)
}

// Retrieves, but does not remove, the head of this queue, or returns null if this queue is empty.
// E peek()
func (pq *PriorityQueue[E]) Peek() E {
	if pq.heap.Len() == 0 {
		var zero E
		return zero
	}
	return pq.heap.items[0]
}

// Removes the specified element from this queue if it is present.
// boolean remove(Object o)
func (pq *PriorityQueue[E]) Remove(item E) bool {
	// when type is comparable
	if reflect.TypeOf(item).Comparable() {
		for i, v := range pq.heap.items {
			if reflect.ValueOf(v).Interface() == reflect.ValueOf(item).Interface() {
				heap.Remove(pq.heap, i)
				return true
			}
		}
		return false
	}

	// panic if not comparable and equals function is not provided
	if pq.equals == nil {
		panic("Type is not comparable and equals function is not provided")
	}

	// use equals function
	for i, v := range pq.heap.items {
		if pq.equals(v, item) {
			heap.Remove(pq.heap, i)
			return true
		}
	}
	return false
}

// Returns the number of elements in this queue.
// int size()
func (pq *PriorityQueue[E]) Size() int {
	return len(pq.heap.items)
}

// Returns an array containing all of the elements in this queue.
// Object[] toArray()
func (pq *PriorityQueue[E]) ToArray() []E {
	return append([]E(nil), pq.heap.items...)
}

// internalHeap is an internal type that implements heap.Interface.
type internalHeap[E any] struct {
	items      []E
	comparator util.Comparator[E]
}

// Len is the number of elements in the collection.
// It is used by the heap package.
func (ph internalHeap[E]) Len() int {
	return len(ph.items)
}

// Less reports whether the element with index i should sort before the element with index j.
// It is used by the heap package.
func (ph internalHeap[E]) Less(i, j int) bool {
	return ph.comparator(ph.items[i], ph.items[j]) < 0
}

// Swap swaps the elements with indexes i and j.
// It is used by the heap package.
func (ph *internalHeap[E]) Swap(i, j int) {
	ph.items[i], ph.items[j] = ph.items[j], ph.items[i]
}

// Push pushes the element x onto the heap.
// It is used by the heap package.
func (ph *internalHeap[E]) Push(x any) {
	item, ok := x.(E)
	if !ok {
		return
	}
	ph.items = append(ph.items, item)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// It is used by the heap package.
func (ph *internalHeap[E]) Pop() any {
	old := ph.items
	n := len(old)
	item := old[n-1]
	ph.items = old[0 : n-1]
	return item
}
