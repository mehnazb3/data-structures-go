// Queue
// Works on FIFO
// Enqueue
// Dequeue

package main

import "fmt"

type Queue []int

func (q *Queue) enQueue(data int) int {
	*q = append(*q, data)
	return data
}

func (q *Queue) deQueue() {
	*q = (*q)[1:]
	return
}

func main() {
	var myQueue Queue
	myQueue.enQueue(10)
	myQueue.enQueue(20)
	myQueue.enQueue(30)
	fmt.Println("My Queue:", myQueue)
	myQueue.deQueue()
	myQueue.deQueue()
	myQueue.deQueue()
	fmt.Println("My Queue after deletion:", myQueue)
}
