package octopus

const (
	defaultMaxDepth int16 = 2
	anchorTag             = "a"
	anchorAttrb           = "href"
)

// New - Creates an Instance of the Octopus Crawler with the given options.
func New(opt CrawlOptions) *webOctopus {
	oct := &webOctopus{
		CrawlOptions: opt,
		visited:      nil,
	}
	oct.setup()
	return oct
}
