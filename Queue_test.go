package Queue

import "testing"

var q Queue

func TestNewQueue(t *testing.T) {
	q := NewQueue(1)

	if len(q.items) != 1 {
		t.Error("Error in size of queue")
	}

	if cap(q.items) != 1 {
		t.Error("Error with capacity of queue")
	}
}

func TestEnqueue(t *testing.T) {
	q.enqueue(0)
	if q.items[0] != 0 {
		t.Error("Error enqueueing item")
	}

	q.enqueue(1)
	if q.items[1] != 1{
		t.Error("Error enqueueing item")
	}
}

func TestPeek(t *testing.T) {
	q.enqueue(0)
	x := q.peek()
	if x != 0 {
		t.Error("Error retrieving last element added to queue")
	}
}

func TestDequeue(t *testing.T) {
	q.enqueue(0)
	x := q.dequeue()
	if x != 0 {
		t.Error("Error retrieving dequeue")
	}
}