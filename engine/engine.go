package engine

import (
	"github.com/zombiecd/zhenai_project/fetch"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]

		requests = requests[1:]

		log.Printf("fetching %s", r.Url)
		body, err := fetch.Fetch(r.Url)

		if err != nil {
			log.Printf("fetcher:error"+"fetcher.url %s,%v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
		log.Printf("go get item %v", item)
		}

	}
}
