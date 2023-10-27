package service

import (
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/repository"
)

type CatService interface {
	AddNewCat(addCatDto *model.AddCatDto) (*entity.Cat, error)
	GetById(id int) (*entity.Cat, error)
	GetAll() ([]entity.Cat, error)
}

type catServiceImpl struct {
	CatRepo repository.CatRepository
}

func (c catServiceImpl) GetAll() ([]entity.Cat, error) {
	return c.CatRepo.FindAll()
}

func (c catServiceImpl) GetById(id int) (*entity.Cat, error) {
	return c.CatRepo.FindById(id)
}

type CatServiceCfg struct {
	CatRepo repository.CatRepository
}

func (c catServiceImpl) AddNewCat(addCatDto *model.AddCatDto) (*entity.Cat, error) {
	cat := &entity.Cat{
		Name:    addCatDto.Name,
		Breed:   addCatDto.Breed,
		Color:   addCatDto.Color,
		BirthAt: addCatDto.BirthAt,
	}

	rs, err := c.CatRepo.Save(cat)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func NewCatService(cfg *CatServiceCfg) CatService {
	return &catServiceImpl{
		CatRepo: cfg.CatRepo,
	}
}
