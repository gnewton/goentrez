package goentrez

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type ESearchQuery struct {
	DB         string
	DateType   string
	Field      string
	MaxDate    string
	MinDate    string
	RelDate    string
	RetMax     uint64
	RetStart   uint64
	Sort       string
	Term       string
	UseHistory bool
	WebEnv     string
}

const baseUrl = "http://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?"

const esearchUrl = baseUrl + "db=pubmed&term=cancer&reldate=60&datetype=edat&retmax=100&usehistory=y"

var transport *http.Transport

func init() {
	transport = &http.Transport{
		DisableKeepAlives:  false,
		DisableCompression: false,
	}
}

// URL arguments, not including leading ?
//
func (q *ESearchQuery) ToUrlQuery() (string, error) {
	err := q.validate()
	if err != nil {
		return "", err
	}
	query := ""
	query = add(query, PARAM_DB, q.DB, true)
	query = add(query, PARAM_DATETYPE, q.DateType, false)
	query = add(query, PARAM_MAXDATE, q.MaxDate, false)
	query = add(query, PARAM_MINDATE, q.MinDate, false)
	//query = add(query, PARAM_QUERY_KEY, q.QueryKey, false)
	query = add(query, PARAM_RELDATE, q.RelDate, false)
	if q.RetMax == 0 {
		q.RetMax = 5
	}
	query = addInt(query, PARAM_RETMAX, q.RetMax, false)
	query = addInt(query, PARAM_RETSTART, q.RetStart, false)
	query = add(query, PARAM_TERM, q.Term, false)
	query = addBool(query, PARAM_USE_HISTORY, q.UseHistory, false)
	query = add(query, PARAM_WEB_ENV, q.WebEnv, false)
	return query, nil
}
func (q *ESearchQuery) validate() error {
	if len(q.DB) != 0 {
		if _, ok := VALID_DATABASES[q.DB]; !ok {
			return errors.New("Unknown database:[" + q.DB + "]")
		}
	}
	return nil
}

func (q *ESearchQuery) getResults() (*EeSearchResult, error) {
	client := &http.Client{Transport: transport}
	var urlParams string
	urlParams, err := q.ToUrlQuery()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("ESearchQuery.Search: " + baseUrl + urlParams)
	resp, err := client.Get(baseUrl + urlParams)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result EeSearchResult
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = xml.Unmarshal(bytes, &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Count=" + result.ECount.Text)
	if result.EWebEnv != nil {
		log.Println("WebEnv=" + result.EWebEnv.Text)
	}
	return &result, nil

}

func pushIdsToChannel(idList []*EId, ids chan string) {
	for i := 0; i < len(idList); i++ {
		log.Println("Pushing id:" + idList[i].Text)
		ids <- idList[i].Text
	}
}

func (q *ESearchQuery) search(ids chan string, quit chan bool) (uint64, *ETranslationSet, *ETranslationStack, error) {
	var count uint64

	results, err := q.getResults()
	if err != nil {
		return 0, nil, nil, err
	}

	if results.EWebEnv != nil {
		q.WebEnv = results.EWebEnv.Text
	}

	pushIdsToChannel(results.EIdList.EId, ids)
	log.Println("++++++++++")
	log.Println(len(ids))
	log.Println("++++++++++")

	count, err = strconv.ParseUint(results.ECount.Text, 10, 64)
	if err != nil {
		return 0, nil, nil, err
	}

	go func() {
		log.Println("###############################")
		for {
			select {
			case <-quit:
				close(ids)
				return
			default:
				log.Println("++++++++++")
				log.Println(len(ids))
				log.Println("++++++++++")
				if len(ids) < 5 {
					q.RetStart = q.RetStart + q.RetMax
					if q.RetStart >= count {
						close(ids)
						return
					}
					results, _ = q.getResults()
					pushIdsToChannel(results.EIdList.EId, ids)
					fmt.Println(results)
				}
			}
		}
	}()

	return count, nil, nil, nil
}

func (q *ESearchQuery) Search() (uint64, chan string, *ETranslationSet, *ETranslationStack, error) {
	//ids := make(chan (string), q.RetMax)
	ids := make(chan (string), 11)
	quit := make(chan (bool))

	count, translationSet, translationStack, err := q.search(ids, quit)

	//for id := range ids {
	for i := 0; i < 55; i++ {
		id := <-ids
		log.Println(i)
		log.Println("------------- " + id)
	}
	quit <- true
	return count, ids, translationSet, translationStack, err

	//return ids, count, translationSet, translationStack, err
}

func addBool(q string, k string, v bool, first bool) string {
	q = insertSeparator(q)
	val := "n"
	if v {
		val = "y"
	}
	return q + k + "=" + val
}

func addInt(q string, k string, v uint64, first bool) string {
	if v == 0 {
		return q
	}
	q = insertSeparator(q)
	return q + k + "=" + strconv.FormatUint(v, 10)
}

func add(q string, k string, v string, first bool) string {
	if len(v) == 0 {
		return q
	}

	q = insertSeparator(q)
	return q + k + "=" + v

}

func insertSeparator(q string) string {
	if len(q) > 0 {
		q = q + "&"
	}
	return q
}
