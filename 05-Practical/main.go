package main

import (
	"fmt"
	"sync"
	"time"
)

type UserDetails struct {
	UserID   int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func main() {
	startTime := time.Now()
	getUserProfile(10)

	fmt.Println("Time Taken: ", time.Since(startTime))
}

func getUserProfile(UserID int) {
	wg := &sync.WaitGroup{}
	resCh := make(chan Response, 3)

	go getComments(UserID, resCh, wg)
	go getLikes(UserID, resCh, wg)
	go getFriends(UserID, resCh, wg)

	wg.Add(3)
	wg.Wait()
	close(resCh)

	for res := range resCh {
		fmt.Println(res.data)
	}

}

func getComments(userID int, resCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{"This is awesome", "Congratulations"}

	resCh <- Response{
		data: comments,
		err:  nil,
	}
	wg.Done()
}

func getLikes(userID int, resCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	likes := 15

	resCh <- Response{
		data: likes,
		err:  nil,
	}
	wg.Done()
}

func getFriends(userID int, resCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	friends := []int{743, 65, 853}

	resCh <- Response{
		data: friends,
		err:  nil,
	}
	wg.Done()
}
