package factories

import (
	"github.com/elastic/go-elasticsearch/v6"
)

// CreateClient creates a new client.
func CreateClient() (client *elasticsearch.Client, err error) {
	client, err = elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{
				"http://elasticsearch:9200",
			},
		},
	)
	return
}
