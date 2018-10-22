package experimental

import (
	"fmt"
)

func Temp() {
	fmt.Println("Hello from experimental section")
}

func Spit(s string) {
	fmt.Printf("%s\n", s)
}

func NewEngine() *spider {
	newSpider := &spider{
		ingestPipeline: nil,
	}
	newSpider.startEngine(true)
	return newSpider
}

func (s *spider) Consume(urlString string) {
	if s.IsRunning() {
		s.ingestPipeline <- urlString
	}
}

func (s *spider) IsRunning() bool {
	return s.isRunning
}

func (s *spider) TurnOff() bool {
	if s.IsRunning() {
		s.quitChannel <- true
		return true
	}
	return false
}

func (s *spider) RestartEngine() bool {
	if !s.IsRunning() {
		s.startEngine(false)
		return true
	}
	return false
}

