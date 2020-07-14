package Queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	//入队
	EnQueue(interface{})
	//出队
	Dequeue() interface{}
	GetFront() interface{}
}
