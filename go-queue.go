package go_queue

import "fmt"

type Queue struct {
	size     int
	counter  int
	elements []any
}

type LinkedListQueue struct {
	data  any
	next  *LinkedListQueue
	first *LinkedListQueue
	last  *LinkedListQueue
}

func NewLinkedListQueue(data int, next *LinkedListQueue) *LinkedListQueue {
	return &LinkedListQueue{data: data, next: next, first: nil, last: nil}
}
func (list *LinkedListQueue) EnQueueLinkedList(data any) *LinkedListQueue {

	if list.first == nil {
		newList := &LinkedListQueue{
			data: data,
			next: nil,
		}
		newList.first = newList
		newList.last = newList
		list = newList
		return list
	} else if list.first == list.last {
		listVal := &LinkedListQueue{}
		listVal.data = data
		listVal.next = nil
		list.next = listVal
		list.last = listVal
		return list
	} else {
		lastList := &LinkedListQueue{}
		lastList.data = data
		lastList.next = nil
		list.last.next = lastList
		list.last = lastList
		return list
	}
}
func NewQueue(size int) *Queue {
	elements := make([]any, size)
	counter := -1
	return &Queue{size, counter, elements}
}
func (q *Queue) EnQueue(element any) {
	if q.size > q.counter {
		q.counter = q.counter + 1
		q.elements[q.counter] = element
	}

}
func (q *Queue) DeQueue() {
	if q.counter != -1 {
		q.elements = q.elements[1:]
	}
}
func (q *Queue) GetLength() int {
	return q.counter
}
func (q *Queue) Display() {
	for i := 0; i <= q.GetLength(); i++ {
		fmt.Println(q.elements[i])
	}
}
func (q *Queue) IsEmpty() bool {
	if q.counter == -1 {
		return true
	}
	return false
}
func (q *Queue) Peek(pos int) any {
	if pos < q.counter {
		return q.elements[pos]
	} else {
		return -1
	}
}
