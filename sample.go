package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v6"

	"go-es-sample/elasticsearch/factories"
	"go-es-sample/elasticsearch/queries"
	"go-es-sample/elasticsearch/requests"
	"go-es-sample/elasticsearch/responses"
	"go-es-sample/elasticsearch/sources"
)

func main() {
	client, _ := factories.CreateClient()
	log.Println(elasticsearch.Version)
	log.Println(client.Info())

	query(client)
	//multiQuery(client)
}

func query(client *elasticsearch.Client) {
	q := queries.SearchQuery{
		Query: (&queries.MatchAll{}).GetQuery(),
	}
	buf, _ := queries.MakeSearchQuery(&q)
	log.Printf("%v\n", buf)

	esRes, _ := requests.Search(client, buf, &requests.SearchConfig{Index: "photo"})
	defer func() {
		_ = esRes.Body.Close()
	}()
	log.Printf("%v\n", esRes)

	res, _ := responses.NewSearch(esRes)
	log.Printf("%v\n", res)
}

func multiQuery(client *elasticsearch.Client) {
	q := []queries.MultiSearchQuery{
		{
			Header: queries.MultiSearchQueryHeader{
				Index: "photo",
			},
			Body: queries.MultiSearchQueryBody{
				Query: (&queries.MatchAll{}).GetQuery(),
				From:  0,
				Size:  10,
			},
		},
	}
	buf, _ := queries.MakeMultiSearchQuery(&q)
	log.Printf("%v\n", buf.String())

	esRes, _ := requests.MultiSearch(
		&requests.SearchConfig{
			Client: client,
			Query:  &buf,
			Index:  "photo",
		},
	)
	defer func() {
		_ = esRes.Body.Close()
	}()
	log.Printf("%v\n", esRes)

	res, _ := responses.NewMultiSearch(esRes)
	log.Printf("%v\n", res)

	n := len(res.Responses)
	for i := 0; i < n; i++ {
		photos := sources.NewPhotoSources(&res.Responses[i])
		log.Printf("photos: %v\n", photos)
	}
}
