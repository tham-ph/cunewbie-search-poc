package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/tham-ph/cunewbie-search-poc/search/src/database"
	"log"
	"strings"
)

func main() {
	es := database.ConnectElasticSearch()

	//res, err := es.Create("jobs", "2", esutil.NewJSONReader(`{"name": "test4324", "what": "Gi432423"}`))
	//res, err := es.Index(
	//	"test",                                  // Index name
	//	strings.NewReader(`{"title" : "Test"}`), // Document body
	//	es.Index.WithDocumentID("1"),            // Document ID
	//	es.Index.WithRefresh("true"),            // Refresh
	//)

	req := esapi.IndexRequest{
		Index:      "students",
		Body:       strings.NewReader(`{"name": "Poom5", "age": 32}`),
		DocumentID: "4",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
