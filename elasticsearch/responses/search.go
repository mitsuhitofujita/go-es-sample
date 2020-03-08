package responses

import (
	"encoding/json"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type SearchSource map[string]interface{}

func (src SearchSource) GetInt(k string, d int) int {
	v, isExists := src[k]
	if !isExists {
		return d
	}
	i, isInt := v.(int)
	if !isInt {
		return d
	}
	return i
}

func (src SearchSource) GetStr(k, d string) string {
	v, isExists := src[k]
	if !isExists {
		return d
	}
	s, isStr := v.(string)
	if !isStr {
		return d
	}
	return s
}

type SearchErrorRootCause struct {
	Type         string `json:"type"`
	Reason       string `json:"reason"`
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	IndexUUID    string `json:"index_uuid"`
	Index        string `json:"index"`
}

type SearchError struct {
	RootCause    []SearchErrorRootCause `json:"root_cause"`
	Type         string                 `json:"type"`
	Reason       string                 `json:"reason"`
	ResourceType string                 `json:"resource_type"`
	ResourceID   string                 `json:"resource_id"`
	IndexUUID    string                 `json:"index_uuid"`
	Index        string                 `json:"index"`
}

type SearchShard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type SearchHitsHit struct {
	Index  string       `json:"_index"`
	Type   string       `json:"_type"`
	ID     string       `json:"_id"`
	Score  float32      `json:"_score"`
	Source SearchSource `json:"_source"`
}

type SearchHits struct {
	Total    int             `json:"total"`
	MaxScore float32         `json:"max_score"`
	Hits     []SearchHitsHit `json:"hits"`
}

type Search struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   SearchShard `json:"_shards"`
	Hits     SearchHits  `json:"hits"`
	Error    SearchError `json:"error"`
	Status   int         `json:"status"`
}

type MultiSearch struct {
	Responses []Search `json:"responses"`
}

func NewSearch(esRes *esapi.Response) (res *Search, err error) {
	res = &Search{}
	var body []byte
	body, err = ioutil.ReadAll(esRes.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, res)
	return
}

func NewMultiSearch(esRes *esapi.Response) (res *MultiSearch, err error) {
	res = &MultiSearch{}
	var body []byte
	body, err = ioutil.ReadAll(esRes.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, res)
	return
}

func (ms *MultiSearch) EachResponse(f func(res *Search, i int)) {
	n := len(ms.Responses)
	for i := 0; i < n; i++ {
		f(&ms.Responses[i], i)
	}
}

func (s *Search) EachSource(f func(src *SearchSource, i int)) {
	s.Hits.EachSource(f)
}

func (s *SearchHits) EachSource(f func(src *SearchSource, i int)) {
	n := len(s.Hits)
	for i := 0; i < n; i++ {
		f(&s.Hits[i].Source, i)
	}
}
