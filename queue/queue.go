package queue

// https://marketsplash.com/golang-queue/

// not thread safe implementation
type Queue[T any] []T

// not thread safe implementation
func (q *Queue[T]) Enqueue(item T) {
	*q = append(*q, item)
}

// Dequeue item is returned
//
// not thread safe implementation
func (q *Queue[T]) Dequeue() T {
	result := (*q)[0]
	*q = (*q)[1:]

	return result
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}
