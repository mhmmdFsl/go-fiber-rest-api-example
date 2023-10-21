package repository

import (
	"fmt"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/httperror"
	"gorm.io/gorm"
)

type CatRepository interface {
	Save(cat *entity.Cat) (*entity.Cat, error)
	FindById(id int) (*entity.Cat, error)
}

type catRepoSitoryImpl struct {
	Db *gorm.DB
}

func (c *catRepoSitoryImpl) FindById(id int) (*entity.Cat, error) {
	var cat entity.Cat
	result := c.Db.First(&cat, id)
	if err := result.Error; err != nil {
		if result.RowsAffected == 0 {
			return nil, httperror.NewFailed(fmt.Sprintf("cat with id %v not found", id), httperror.NOT_FOUND)
		}
		return nil, httperror.NewFailed(err.Error(), httperror.INTERNAL)
	}
	return &cat, nil
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
