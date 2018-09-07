package engine

import (
	"crawer/fetcher"
	"log"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests,r)
	}

	for len(requests)>0 {
		r := requests[0]
		requests = requests[1:]

		//log.Printf("Fetching %s\n", r.Url)
		//body, e := fetcher.Fetch(r.Url)
		//if e != nil{
		//	log.Printf("Fetcher: error fetching url %s: %v",
		//		r.Url, e)
		//	continue
		//}
		//
		//parseResult := r.PaserFunc(body)

		parseResult, e := worker(r)
		if e != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item :=range parseResult.Items{
			//%v 不转义；
			log.Printf("Got item %v", item)
		}
	}
}

func worker( r Request) (ParseResult,error) {
	log.Printf("Fetching %s\n", r.Url)
	body, e := fetcher.Fetch(r.Url)
	if e != nil{
		log.Printf("Fetcher: error fetching url %s: %v",
			r.Url, e)
		return ParseResult{},e
	}

	return r.PaserFunc(body), nil
}