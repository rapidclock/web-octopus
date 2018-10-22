package experimental

import "time"

type spider struct {
	ingestPipeline chan<- string
	quitChannel    chan<- bool
	isRunning      bool
}

func (s *spider) indicateOn() {
	s.isRunning = true
}

func (s *spider) indicateOff() {
	s.isRunning = false
}

func (s *spider) startEngine(enableDelay bool) {
	ingestChannel := make(chan string)
	quitCh := make(chan bool)
	s.ingestPipeline = ingestChannel
	s.quitChannel = quitCh
	go func() {
		s.indicateOn()
		defer close(ingestChannel)
		defer close(quitCh)
		defer s.indicateOff()
		for {
			select {
			case urlStr := <-ingestChannel:
				Spit(urlStr)
			case <-quitCh:
				return
			}
		}
	}()
	if enableDelay {
		<-time.After(500 * time.Millisecond)
	}
}

