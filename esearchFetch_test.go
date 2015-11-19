package goentrez

import (
	"fmt"
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	q := new(ESearchFetch)
	//m := []string{"2344", "44511", "654125"}
	m := []string{"zzzzzzzzzzzz"}

	//records, err := q.FetchList(m)
	arts, bookArts, err := q.FetchList(m)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(arts, bookArts, err)
	}

}
