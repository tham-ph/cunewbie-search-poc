package database

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElasticSearch() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "123456",
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return es
}
