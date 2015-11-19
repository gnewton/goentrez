package goentrez

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const fetchBaseUrl = "http://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?"

type ESearchFetch struct {
}

func (q *ESearchFetch) FetchList(ids []string) ([]*Chi_PubmedArticle, []*Chi_PubmedBookArticle, error) {
	idString := ""
	for _, id := range ids {
		if idString != "" {
			idString = idString + ","
		}
		idString = idString + id
	}
	url := fetchBaseUrl + "db=pubmed&retmode=xml&id=" + idString
	log.Println("999222888 " + url)

	client := &http.Client{Transport: transport}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	var result Chi_PubmedArticleSet
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	var fe EFetchResultError
	err = xml.Unmarshal(bytes, &fe)
	log.Println(fe)
	//log.Println(errorRoot.FetchError)
	if fe.Error != nil {
		log.Println("***************************")
		return nil, nil, errors.New("Error in request: Error message from eutils: [" + fe.Error.Text + "]  url=" + url)
	}
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	err = xml.Unmarshal(bytes, &result)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	for _, r := range result.Chi_PubmedArticle {
		log.Println(r.Chi_MedlineCitation.Chi_DateCompleted.Chi_Year.Text)
	}
	return result.Chi_PubmedArticle, result.Chi_PubmedBookArticle, nil
}
