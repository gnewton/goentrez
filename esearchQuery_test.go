package goentrez

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	//q := ESearchQuery{DB: "pubmed", Term: "cancer", UseHistory: true}
	q := ESearchQuery{Term: "cancer"}

	count, ids, translationSet, translationStack, err := q.Search()
	log.Println(ids)
	log.Println(count)
	log.Println(translationSet)
	log.Println(translationStack)
	log.Println(err)

	//for id := range ids {
	for i := 0; i < 2055; i++ {
		id := <-ids
		log.Println(i)
		log.Println("------------- " + id)
	}
	q.Done()

}
