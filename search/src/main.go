package main

import (
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/tham-ph/cunewbie-search-poc/search/src/database"
	"log"
)

func main() {
	es := database.ConnectElasticSearch()

	res, err := es.Create("students", "14", esutil.NewJSONReader(map[string]interface{}{"name": "Poom haha", "age": 33}))
	res, err = es.Create("students", "13", esutil.NewJSONReader(map[string]interface{}{"name": "Poom zaza", "age": 33}))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"message": map[string]interface{}{
					"query": "Poom",
				},
			},
		},
	}
	res, err = es.Search(es.Search.WithIndex("students"), es.Search.WithBody(esutil.NewJSONReader(query)))
	log.Println(res)
}
