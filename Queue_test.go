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
	q.Enqueue(0)
	if q.items[0] != 0 {
		t.Error("Error enqueueing item")
	}

	q.Enqueue(1)
	if q.items[1] != 1{
		t.Error("Error enqueueing item")
	}
}

func TestPeek(t *testing.T) {
	q.Enqueue(0)
	x := q.Peek()
	if x != 0 {
		t.Error("Error retrieving last element added to queue")
	}
}

func TestDequeue(t *testing.T) {
	q.Enqueue(0)
	x := q.Dequeue()
	if x != 0 {
		t.Error("Error retrieving dequeue")
	}
}

func TestQueue_IsEmpty(t *testing.T) {

	q.Enqueue(0)
	x := q.IsEmpty()
	if x {
		t.Error("Queue should have a value")
	}
}