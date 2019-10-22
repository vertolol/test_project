package elastic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
)


func (client ElasticWorker) GetIdsByName(name string) (results []int64, err error){
	matchQuery := elastic.NewMatchQuery("name", name).Operator("AND")
	searchResult, err := client.Client.Search().
		Index(client.Index).
		Query(matchQuery).
		Do(context.Background())
	if err != nil {
		return results, err
	}


	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			id, err := strconv.ParseInt(hit.Id, 10, 64)
			if err != nil {
				return results, err
			}
			results = append(results, id)
		}
	} else {
		fmt.Printf("Found no %v\n", client.Index)
	}

	return results, err
}
