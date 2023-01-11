package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	userID := 10
	respch := make(chan string, 100)

	wg := &sync.WaitGroup{}

	go fetchUserData(userID, respch, wg)
	go fetchUserRecommendations(userID, respch, wg)
	go fetchUserLikes(userID, respch, wg)
	wg.Add(3)
	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(now))

}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(80 * time.Millisecond)
	respch <- "User Data"
}

func fetchUserRecommendations(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)

	respch <- "User Recommendations"
}

func fetchUserLikes(userID int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	respch <- "User Likes"
}
