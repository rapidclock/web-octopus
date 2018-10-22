package main

import (
	"fmt"
	exp "github.com/rapidclock/web-octopus/experimental"
	"time"
)

const (
	HomeUrl = "https://en.wikipedia.org/wiki/Main_Page"
)

func main() {
	exp.Temp()
	crawler := exp.NewEngine()
	crawler.Consume(HomeUrl)
	crawler.Consume("Chicken Nugget")
	crawler.Consume("Tina kicks ass!!")
	fmt.Printf("1. %v\n", crawler.IsRunning())
	crawler.TurnOff()
	<-time.After(200 * time.Millisecond)
	fmt.Printf("2. %v\n", crawler.IsRunning())
	crawler.RestartEngine()
	<-time.After(100 * time.Millisecond)
	fmt.Printf("3. %v\n", crawler.IsRunning())
	crawler.Consume("Hello Poppet!!")
	fmt.Printf("4. %v\n", crawler.IsRunning())
	crawler.TurnOff()
}
