package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	newQueue := New(2)
	l := len(newQueue.items)
	c := newQueue.capacity
	fmt.Println(l, c)
	if len(newQueue.items) != 0 && newQueue.capacity != 2 {
		t.Errorf("expected: %d, but got %d\n", l, c)
	}
}

func TestAddQueue(t *testing.T) {
	newQueue := New(3)
	for i := 0; i < 3; i++ {
		if len(newQueue.items) != i {
			t.Errorf("incorrect queue length, expected %v, but go %v\n", i, len(newQueue.items))
		}
		if !newQueue.Append(i) {
			t.Errorf("failed to append item %v to queue", i)
		}
	}
	if newQueue.Append(4) {
		t.Errorf("should not beable to add to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: expected %v, but go %v\n", i, item)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("the queue should be empty, got %v\n", item)
	}
}
