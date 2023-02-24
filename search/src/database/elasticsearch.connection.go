package database

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElasticSearch() *elasticsearch.Client {
	//cfg := elasticsearch.Config{
	//	Addresses: []string{
	//		"https://localhost:9200",
	//	},
	//	Username: "elastic",
	//	Password: "123456",
	//}
	//es, err := elasticsearch.NewClient(cfg)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	return es
}
