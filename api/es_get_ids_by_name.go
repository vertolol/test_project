package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)

func getIdsByName(name string) []int64{
	var (
		indexName = "products"
		results []int64
	)

	client, err := elastic.NewClient(elastic.SetURL("http://elastic:9200"))
	ifPanic(err)

	matchQuery := elastic.NewMatchQuery("name", name).Operator("AND")
	searchResult, err := client.Search().
		Index(indexName).
		Query(matchQuery).
		Do(context.Background())
	ifPanic(err)

	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			id, err := strconv.ParseInt(hit.Id, 10, 64)
			ifPanic(err)
			results = append(results, id)
		}
	} else {
		fmt.Printf("Found no %v\n", indexName)
	}

	return results
}
