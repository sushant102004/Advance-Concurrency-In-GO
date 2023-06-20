package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Important - If a go channel is FULL then it will block the program.
	*/

	resultCh := make(chan string, 10)

	resultCh <- "A"
	resultCh <- "B"
	resultCh <- "C"
	resultCh <- "D"
	close(resultCh)

	for {
		res, ok := <-resultCh
		if !ok {
			break
		}
		fmt.Println(res)
	}

	// for res := range resultCh {
	// 	fmt.Println(res)
	// }

}

func fetchResources(val string) string {
	time.Sleep(time.Second * 2)
	return val
}
