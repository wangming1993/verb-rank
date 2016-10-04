package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	channel()
}

func channel() {
	done := make(chan bool)
	go parse(done)
	product(done)
}

func product(done chan<- bool) {
	fmt.Println("start sleep 2 sec...")
	time.Sleep(2 * time.Second)
	fmt.Println("wake up...")
	done <- true
	time.Sleep(1 * time.Second)
}

func parse(done <-chan bool) {
	var finished bool
	for {
		if finished {
			break
		}
		select {
		case <-done:
			fmt.Println("get from chan done..")
			finished = true
		}
	}
	fmt.Println("finished...")
}
