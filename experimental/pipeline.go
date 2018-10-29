package experimental

func BuildSystem() chan<- string {
	parsePipe, compPipeChan := MakeParsingPipe()
	reqPipe := MakeRequistionPipe(parsePipe, nil)
	validationPipe := MakeUrlValidationPipe(reqPipe)
	unduplPipe := MakeUnduplicationPipe(validationPipe)
	compPipe := MakeCompositionPipe(unduplPipe)
	compPipeChan <- compPipe
	return listenForUrl(compPipe)
}

func listenForUrl(compositionPipe chan<- *ReqProp) chan<- string {
	listenChan := make(chan string)
	go func() {
		for {
			select {
			case urlStr := <-listenChan:
				compositionPipe <- &ReqProp{
					"",
					urlStr,
				}
			}
		}
	}()
	return listenChan
}
