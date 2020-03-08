package sources

import (
	"bytes"
	"context"
	"encoding/json"
	"go-es-sample/elasticsearch/responses"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

func SearchPhoto(client *elasticsearch.Client) (res *esapi.Response, err error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if e := json.NewEncoder(&buf).Encode(query); e != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("photo"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	return
}

type PhotoSource struct {
	src *responses.SearchSource
}

func NewPhotoSources(res *responses.Search) (sources *[]PhotoSource) {
	s := make([]PhotoSource, len(res.Hits.Hits))

	res.EachSource(func(src *responses.SearchSource, i int) {
		s[i] = NewPhotoSource(src)
	})

	sources = &s
	return
}

func NewPhotoSource(src *responses.SearchSource) PhotoSource {
	return PhotoSource{src: src}
}

func (photoSource PhotoSource) GetTier() int {
	return photoSource.src.GetInt("tier", 0)
}

func (photoSource PhotoSource) GetSubject() string {
	return photoSource.src.GetStr("subject", "")
}

func (photoSource PhotoSource) GetTags() []string {
	tags := photoSource.src.GetStr("tags", "")
	return strings.Split(tags, " ")
}
