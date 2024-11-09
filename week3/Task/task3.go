package main

import (
	"fmt"
	"sync"
)

func main() {
	str1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str2 := "0123456789"

	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(str1); i++ {
			fmt.Println(string(str1[i]), i)
			if (i+1)%2 == 0 {
				ch <- struct{}{}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < len(str2); j++ {
			<-ch
			fmt.Println(string(str2[j]), j)
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Println()
}
