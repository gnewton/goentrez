package goentrez

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	q := ESearchQuery{DB: "pubmed", Term: "cancer", UseHistory: true}

	ids, totalCount, translationSet, translationStack, err := q.Search()
	log.Println(ids)
	log.Println(totalCount)
	log.Println(translationSet)
	log.Println(translationStack)
	log.Println(err)

}
