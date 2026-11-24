package main

import "fmt"

// Queue implements FIFO on two LIFO stacks: "in" collects pushes, "out"
// serves pops in reversed (i.e. arrival) order, refilled only when empty.
type Queue struct {
	in, out []int
}

// Enqueue pushes onto the in-stack. O(1).
func (q *Queue) Enqueue(v int) {
	q.in = append(q.in, v)
}

// Dequeue pops from the out-stack, first draining the in-stack into it if
// needed. Amortized O(1): each element crosses over exactly once.
func (q *Queue) Dequeue() (int, bool) {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			last := len(q.in) - 1
			q.out = append(q.out, q.in[last])
			q.in = q.in[:last]
		}
	}
	if len(q.out) == 0 {
		return 0, false
	}
	last := len(q.out) - 1
	v := q.out[last]
	q.out = q.out[:last]
	return v, true
}

// Front returns the next value Dequeue would yield.
func (q *Queue) Front() (int, bool) {
	if len(q.out) == 0 {
		if len(q.in) == 0 {
			return 0, false
		}
		return q.in[0], true
	}
	return q.out[len(q.out)-1], true
}

// Size reports how many values are queued.
func (q *Queue) Size() int { return len(q.in) + len(q.out) }

func main() {
	var q Queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	for q.Size() > 0 {
		v, _ := q.Dequeue()
		fmt.Print(v, " ")
	}
	fmt.Println()
}
