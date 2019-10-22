package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)


type ElasticWorker struct {
	client *elastic.Client
}


func createElasticConnection() *elastic.Client{
	url := fmt.Sprintf("http://%s:%s",
		ELASTIC_HOST,
		ELASTIC_PORT,
	)

	client, err := elastic.NewClient(elastic.SetURL(url))
	ifPanic(err)

	return client
}


func (client *ElasticWorker) getIdsByName(name string) []int64{
	matchQuery := elastic.NewMatchQuery("name", name).Operator("AND")
	searchResult, err := client.client.Search().
		Index(ELASTIC_INDEX_NAME).
		Query(matchQuery).
		Do(context.Background())
	ifPanic(err)

	var results []int64
	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			id, err := strconv.ParseInt(hit.Id, 10, 64)
			ifPanic(err)
			results = append(results, id)
		}
	} else {
		fmt.Printf("Found no %v\n", ELASTIC_INDEX_NAME)
	}

	return results
}
