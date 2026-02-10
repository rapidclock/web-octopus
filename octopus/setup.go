package octopus

import (
	"time"

	"golang.org/x/time/rate"
)

func (o *octopus) setupOctopus() {
	o.setupValidProtocolMap()
	o.setupTimeToQuit()
	o.setupMaxLinksCrawled()
	o.setupRateLimiting()
}

func (o *octopus) setupValidProtocolMap() {
	o.isValidProtocol = make(map[string]bool)
	for _, protocol := range o.ValidProtocols {
		o.isValidProtocol[protocol] = true
	}
}

func (o *octopus) setupTimeToQuit() {
	if o.TimeToQuit > 0 {
		o.timeToQuit = time.Duration(o.TimeToQuit) * time.Second
	} else {
		panic("TimeToQuit is not greater than 0")
	}
}

func (o *octopus) setupMaxLinksCrawled() {
	if o.MaxCrawledUrls == 0 {
		panic("MaxCrawledUrls should either be negative or greater than 0.")
	}
}

func (o *octopus) setupRateLimiting() {
	crawlRate := rate.Limit(o.CrawlRatePerSec)
	burstLimit := rate.Limit(o.CrawlBurstLimitPerSec)
	validateRateLimits(crawlRate, burstLimit)
	if crawlRate > 0 {
		o.rateLimiter = rate.NewLimiter(crawlRate, int(burstLimit))
	}
}

func validateRateLimits(crawlRate, burstLimit rate.Limit) {
	if crawlRate == 0 {
		panic("Crawl Rate can never be zero!")
	} else if burstLimit < crawlRate {
		panic("Burst Crawl Rate Limit should be greater than or equal to the" +
			" Crawl Rate Limit")
	} else if crawlRate < 0 && burstLimit > 0 {
		panic("Cannot set Burst Limit positive when the crawl Rate is not" +
			" greater than zero.")
	}
}
