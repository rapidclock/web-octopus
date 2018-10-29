package experimental

func NewMonster() *Monster {
	listenChan := make(chan string)
	return &Monster{
		listenChan,
		nil,
	}
}

func (m *Monster) BuildSystem(opAdapterPipe chan<- *Node) {
	parsePipe, compPipeChan := MakeParsingPipe()
	var reqPipe chan<- *Node
	if opAdapterPipe == nil {
		reqPipe = MakeRequistionPipe(parsePipe, nil)
	} else {
		reqPipe = MakeRequistionPipe(parsePipe, opAdapterPipe)
	}
	validationPipe := MakeUrlValidationPipe(reqPipe)
	unduplPipe := MakeUnduplicationPipe(validationPipe)
	compPipe := MakeCompositionPipe(unduplPipe)
	compPipeChan <- compPipe
	m.compPipe = compPipe
}

func (m *Monster) StartCrawling(baseUrlString string) {
	for {
		select {
		case urlStr := <-m.listenPipe:
			m.compPipe <- &ReqProp{
				"",
				urlStr,
			}
		}
	}
}
