package elastic

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

type ElasticWorker struct {
	Client 		*elastic.Client
	Index 		string
}

type ConnectionConfig struct {
	HOST 		string
	PORT 		string
	INDEX_NAME 	string
}

func CreateElasticConnection(config *ConnectionConfig) (client *elastic.Client, err error){
	url := fmt.Sprintf("http://%s:%s",
		config.HOST,
		config.PORT,
	)

	client, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return client, err
	}
	return client, err
}
