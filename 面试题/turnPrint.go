package 面试题

import (
	"fmt"
	"strings"
	"sync"
)

/**
交替打印数字和字母
问题描述

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func turnPrint() {
	letter, number := make(chan int), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i:=1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- i
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	go func(group *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					group.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
}
