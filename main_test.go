package main

import "testing"

func TestQueueUsingStacks(t *testing.T) {
	var q Queue
	if _, ok := q.Dequeue(); ok {
		t.Error("dequeue on empty must fail")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	if v, _ := q.Front(); v != 1 {
		t.Errorf("Front = %d, want 1", v)
	}
	if v, _ := q.Dequeue(); v != 1 {
		t.Errorf("Dequeue = %d, want 1 (FIFO)", v)
	}
	// Enqueue after a partial drain must not jump the line.
	q.Enqueue(3)
	if v, _ := q.Dequeue(); v != 2 {
		t.Errorf("Dequeue = %d, want 2", v)
	}
	if v, _ := q.Dequeue(); v != 3 {
		t.Errorf("Dequeue = %d, want 3", v)
	}
	if q.Size() != 0 {
		t.Errorf("Size = %d, want 0", q.Size())
	}
	// Front works when everything still sits on the in-stack.
	q.Enqueue(9)
	if v, ok := q.Front(); !ok || v != 9 {
		t.Errorf("Front = (%d,%v), want 9", v, ok)
	}
}
