package repositories

import "go-es-sample/entities"

type PhotoRepositoryInterface interface {
	All() []entities.PhotoEntityGetterInterface
	Search() []entities.PhotoEntityGetterInterface
}
