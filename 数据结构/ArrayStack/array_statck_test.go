package ArrayStack

import (
	"fmt"
	"log"
	"testing"
)

func TestArrayStack(t *testing.T) {
	stack := Constructor(10)
	for i := 0; i < 10; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}

const (
	a         = iota
	b float32 = iota
	c
)

func TestCount(t *testing.T) {

	fmt.Printf("%v,%T \n", a, a)
	fmt.Printf("%v,%T \n", b, b)
	fmt.Printf("%v,%T \n", c, c)
}

func TestPoint(t *testing.T) {
	switch x := 4; x {
	case 4:
		x += 10
		fmt.Println(x)
		fallthrough
	case 5:
		x += 20
		fmt.Println(x)
		fallthrough
	default:
		fmt.Println(x)

	}
}

func TestReco(t *testing.T){
	defer func(){
		if err := recover();err != nil{
			log.Fatalln(err)
		}
	}()
	panic("i am dead")
	fmt.Println(1)
}

