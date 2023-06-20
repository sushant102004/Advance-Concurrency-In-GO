package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	userID, err := getUserID()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User ID: ", userID)
}

func getUserID() (string, error) {
	// This context will wait till 100 miliseconds and then cancel this function.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	type result struct {
		userID string
		err    error
	}

	resultCh := make(chan result, 1)

	go func() {
		res, err := httpCall()
		resultCh <- result{
			userID: res,
			err:    err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultCh:
		return res.userID, nil
	}

}

func httpCall() (string, error) {
	// This function will take 500 ms to respond
	time.Sleep(time.Millisecond * 500)
	return "User ID: 1", nil
}
