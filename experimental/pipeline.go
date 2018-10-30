package experimental

func NewMonsterWithOptions(options *Options) *Monster {
	listenChan := make(chan string)
	var opt *Options
	if options != nil {
		opt = options
	} else {
		opt = &Options{
			-1,
		}
	}
	return &Monster{
		Options:    opt,
		listenPipe: listenChan,
		compPipe:   nil,
	}
}

func NewMonster() *Monster {
	listenChan := make(chan string)
	return &Monster{
		Options:    nil,
		listenPipe: listenChan,
		compPipe:   nil,
	}
}

func (m *Monster) BuildSystem(opAdapterPipe chan<- *Node) {
	parsePipe, compPipeChan := m.MakeParsingPipe()
	var reqPipe chan<- *Node
	if opAdapterPipe == nil {
		reqPipe = m.MakeRequisitionPipe(parsePipe, nil)
	} else {
		reqPipe = m.MakeRequisitionPipe(parsePipe, opAdapterPipe)
	}
	validationPipe := m.MakeUrlValidationPipe(reqPipe)
	unduplPipe := m.MakeUnduplicationPipe(validationPipe)
	cleanPipe := m.MakeLinkCleaningPipe(unduplPipe)
	compPipe := m.MakeCompositionPipe(cleanPipe)
	compPipeChan <- compPipe
	m.compPipe = compPipe
}

func (m *Monster) StartCrawling(baseUrlString string) {
	go func() {
		m.listenPipe <- baseUrlString
	}()
	for {
		select {
		case urlStr := <-m.listenPipe:
			m.compPipe <- &ReqProp{
				"",
				urlStr,
				0,
			}
		}
	}
}
