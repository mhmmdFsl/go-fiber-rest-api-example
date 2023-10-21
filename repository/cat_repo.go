package repository

import (
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"gorm.io/gorm"
)

type CatRepository interface {
	Save(cat *entity.Cat) (*entity.Cat, error)
}

type catRepoSitoryImpl struct {
	Db *gorm.DB
}

func (c *catRepoSitoryImpl) Save(cat *entity.Cat) (*entity.Cat, error) {

	result := c.Db.Create(cat)
	if err := result.Error; err != nil {
		return nil, err
	}

	return cat, nil
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepoSitoryImpl{
		Db: db,
	}
}
