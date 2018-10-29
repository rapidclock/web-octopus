package octopus

const (
	defaultMaxDepth int16 = 2
	anchorTag             = "a"
	anchorAttrb           = "href"
)

// MakeNew - Creates an Instance of the Octopus Crawler with the given options.
func MakeNew(opt CrawlOptions) *octopus {
	oct := &octopus{
		CrawlOptions: opt,
		visited:      make(map[Node]bool),
	}
	oct.setup()
	return oct
}
