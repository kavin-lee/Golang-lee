package engine

import (
	"Golang-lee/learngo/crwaler/fetcher"
	"log"
)

func Run(sends ...Request) {
	var requests []Request
	for _, r := range sends {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error"+
				"fetching url %s:%v", r.Url, err)
			continue
		}
		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)
		for _, item := range ParseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
