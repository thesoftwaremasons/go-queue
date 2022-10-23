package go_queue

import (
	"reflect"
	"testing"
)

func TestLinkedListQueue_EnQueueLinkedList(t *testing.T) {
	type fields struct {
		data  any
		next  *LinkedListQueue
		first *LinkedListQueue
		last  *LinkedListQueue
	}
	type args struct {
		data any
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   *LinkedListQueue
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := &LinkedListQueue{
				data:  tt.fields.data,
				next:  tt.fields.next,
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			if got := list.EnQueueLinkedList(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnQueueLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkedListQueue(t *testing.T) {
	type args struct {
		data int
		next *LinkedListQueue
	}
	var tests []struct {
		name string
		args args
		want *LinkedListQueue
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedListQueue(tt.args.data, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedListQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewQueue(t *testing.T) {
	type args struct {
		size int
	}
	var tests []struct {
		name string
		args args
		want *Queue
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_DeQueue(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			q.DeQueue()
		})
	}
}

func TestQueue_Display(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			q.Display()
		})
	}
}

func TestQueue_EnQueue(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	type args struct {
		element any
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			q.EnQueue(tt.args.element)
		})
	}
}

func TestQueue_GetLength(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	var tests []struct {
		name   string
		fields fields
		want   int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			if got := q.GetLength(); got != tt.want {
				t.Errorf("GetLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	var tests []struct {
		name   string
		fields fields
		want   bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type fields struct {
		size     int
		counter  int
		elements []any
	}
	type args struct {
		pos int
	}
	var tests []struct {
		name   string
		fields fields
		args   args
		want   any
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:     tt.fields.size,
				counter:  tt.fields.counter,
				elements: tt.fields.elements,
			}
			if got := q.Peek(tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
