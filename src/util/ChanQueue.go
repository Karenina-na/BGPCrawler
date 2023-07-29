package util

import (
	"fmt"
	"sync"
)

// ChanQueue is a queue
type ChanQueue[T any] struct {
	queue   chan T
	len     int
	head    *T
	rwLock  *sync.RWMutex
	opeLock *sync.Mutex
}

// NewChanQueue
//
//	@Description: create a new queue
//	@param size	: the size of the queue
//	@return *ChanQueue[T]	: the new queue
func NewChanQueue[T any](size int) *ChanQueue[T] {
	return &ChanQueue[T]{
		queue:   make(chan T, size),
		len:     0,
		head:    nil,
		rwLock:  &sync.RWMutex{},
		opeLock: &sync.Mutex{},
	}
}

// Enqueue
//
//	@Description: enqueue an object to the queue
//	@receiver Q	: the queue
//	@param data	: the object to be enqueued
func (Q *ChanQueue[T]) Enqueue(data *T) {
	Q.rwLock.Lock()
	Q.len++
	Q.queue <- *data
	Q.rwLock.Unlock()
}

// Dequeue
//
//	@Description: dequeue an object from the queue
//	@receiver Q	: the queue
//	@return T	: the object to be dequeued
func (Q *ChanQueue[T]) Dequeue() *T {
	Q.rwLock.Lock()
	if Q.len == 0 {
		Q.rwLock.Unlock()
		return nil
	}
	var data *T
	Q.len--
	if Q.head != nil {
		data = Q.head
	} else {
		t := <-Q.queue
		data = &t
	}
	if Q.len > 0 {
		t := <-Q.queue
		Q.head = &t
	}
	Q.rwLock.Unlock()
	return data
}

// Length
//
//	@Description: get the length of the queue
//	@receiver Q	: the queue
//	@return int	: the length of the queue
func (Q *ChanQueue[T]) Length() int {
	Q.rwLock.RLock()
	defer Q.rwLock.RUnlock()
	return Q.len
}

// IsEmpty
//
//	@Description: check if the queue is empty
//	@receiver Q	: the queue
//	@return bool	: true if the queue is empty, otherwise false
func (Q *ChanQueue[T]) IsEmpty() bool {
	Q.rwLock.RLock()
	defer Q.rwLock.RUnlock()
	return Q.len == 0
}

// IsFull
//
//	@Description: check if the queue is full
//	@receiver Q	: the queue
//	@return bool	: true if the queue is full, otherwise false
func (Q *ChanQueue[T]) IsFull() bool {
	Q.rwLock.RLock()
	defer Q.rwLock.RUnlock()
	return Q.len == cap(Q.queue)
}

// Head
//
//	@Description: get the head of the queue
//	@receiver Q	: the queue
//	@return T	: the head of the queue
func (Q *ChanQueue[T]) Head() T {
	Q.rwLock.RLock()
	defer Q.rwLock.RUnlock()
	return *Q.head
}

// Clear
//
//	@Description:	clear the queue
//	@receiver Q	:	the queue
func (Q *ChanQueue[T]) Clear() {
	Q.rwLock.Lock()
	Q.len = 0
	Q.head = nil
	Q.rwLock.Unlock()
}

// Destroy
//
//	@Description: destroy the queue
//	@receiver Q	: the queue
func (Q *ChanQueue[T]) Destroy() {
	Q.rwLock.Lock()
	Q.len = 0
	Q.head = nil
	close(Q.queue)
	Q.rwLock.Unlock()
}

// String
//
//	@Description: get the string of the queue
//	@receiver Q	: the queue
//	@return string	: the string of the queue
func (Q *ChanQueue[T]) String() string {
	Q.rwLock.RLock()
	defer Q.rwLock.RUnlock()
	return fmt.Sprintf("ChanQueue: len=%d, cap=%d, head=%v", Q.len, cap(Q.queue), *Q.head)
}

// Operate
//
//	@Description: operate the queue
//	@receiver Q*ChanQueue[T]	: the queue
func (Q *ChanQueue[T]) Operate(f func()) {
	Q.opeLock.Lock()
	f()
	defer Q.opeLock.Unlock()
}
