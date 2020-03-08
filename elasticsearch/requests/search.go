package requests

import (
	"context"
	"io"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type SearchConfig struct {
	Client *elasticsearch.Client
	Index  string
	Query  io.Reader
}

func Search(client *elasticsearch.Client, query io.Reader, config *SearchConfig) (res *esapi.Response, err error) {
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(config.Index),
		client.Search.WithBody(query),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	return
}

func MultiSearch(config *SearchConfig) (res *esapi.Response, err error) {
	client := config.Client
	res, err = client.Msearch(
		config.Query,
		client.Msearch.WithContext(context.Background()),
		client.Msearch.WithPretty(),
		client.Msearch.WithIndex(config.Index),
	)
	return
}
