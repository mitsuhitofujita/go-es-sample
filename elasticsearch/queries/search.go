package queries

import (
	"bytes"
	"encoding/json"
	"io"
)

type Query map[string]interface{}

type SearchQuery struct {
	Query Query `json:"query"`
}

type MultiSearchQueryHeader struct {
	Index string `json:"index"`
}

type MultiSearchQueryBody struct {
	Query Query `json:"query"`
	From  int   `json:"from"`
	Size  int   `json:"size"`
}

type MultiSearchQuery struct {
	Header MultiSearchQueryHeader
	Body   MultiSearchQueryBody
}

type QueryInterface interface {
	GetQuery() (d map[string]interface{})
}

type MatchAll struct {
}

func MakeSearchQuery(query *SearchQuery) (buf io.Reader, err error) {
	var b []byte
	b, err = json.Marshal(query)
	if err != nil {
		return
	}
	buf = bytes.NewReader(b)
	return
}

func MakeMultiSearchQuery(queries *[]MultiSearchQuery) (buf bytes.Buffer, err error) {
	nl := []byte("\n")
	n := len(*queries)
	for i := 0; i < n; i++ {
		var h, b []byte
		h, err = json.Marshal((*queries)[i].Header)
		if err != nil {
			return
		}
		buf.Write(h)
		buf.Write(nl)
		b, err = json.Marshal((*queries)[i].Body)
		if err != nil {
			return
		}
		buf.Write(b)
		buf.Write(nl)
	}
	return
}

func (q *MatchAll) GetQuery() (d map[string]interface{}) {
	d = map[string]interface{}{"match_all": map[string]interface{}{}}
	return
}
