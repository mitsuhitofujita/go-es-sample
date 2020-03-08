package repositories

import (
	"go-es-sample/entities"

	"github.com/elastic/go-elasticsearch/v6"
)

type PhotoRepository struct {
	Client elasticsearch.Client
}

func NewPhotoRepository(client elasticsearch.Client) *PhotoRepository {
	return &PhotoRepository{
		Client: client,
	}
}

func (rep PhotoRepository) All() (photos []entities.PhotoEntity) {

}
