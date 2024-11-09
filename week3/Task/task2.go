package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type num struct {
	ID  int
	Val int
}

func main() {
	nums := make(chan num, 20) //建立channel
	//用等待锁实现20个goroutine的完成
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1) //增加一个等待
		go func(id int) {
			defer wg.Done()               //完成
			randomValue := rand.Intn(100) //随机数
			nums <- num{id, randomValue}
		}(i)
	}
	go func() {
		wg.Wait()
		close(nums)
	}()
	var a []num
	for i := range nums {
		a = append(a, i)
	}
	fmt.Println("排序前的结果:")
	for _, i := range a {
		fmt.Printf("编号：%d   数值：%d\n", i.ID, i.Val)
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].Val < a[j].Val
	})
	fmt.Println("\n排序后的结果:")
	for _, i := range a {
		fmt.Printf("编号：%d   数值：%d\n", i.ID, i.Val)
	}
}
