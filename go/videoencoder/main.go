package main

import (
	"strconv"

	"github.com/renatobrittoaraujo/fullcycle-go-video-encoder/async"
)

var randids = []string{
	"123790-1738921381209",
	"12708317289032178903",
	"123790-1738921381209",
	"12708317289032178903",
	"123790-1738921381209",
	"12708317289032178903",
	"123790-1738921381209",
	"12708317289032178903",
	"123790-1738921381209",
	"12708317289032178903",
}

func main() {
	onlyone := make(chan string)
	done := make(chan string)

	go func() {
		i := 0
		for {
			onlyone <- randids[i%10] + " --- " + strconv.Itoa(i)
			i++
		}
	}()

	for i := 0; i < 10; i++ {
		go async.LongRoutine(onlyone, done)
	}

	<-done
}
