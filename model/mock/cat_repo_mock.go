package mock

import (
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/stretchr/testify/mock"
)

type CatRepoMock struct {
	mock.Mock
}

func (c *CatRepoMock) Save(cat *entity.Cat) (*entity.Cat, error) {

	args := c.Called(cat)
	a0 := args.Get(0)
	a1 := args.Error(1)
	return a0.(*entity.Cat), a1
}

func (c *CatRepoMock) FindById(id int) (*entity.Cat, error) {
	args := c.Called(id)
	a0 := args.Get(0)
	a1 := args.Error(1)
	return a0.(*entity.Cat), a1
}

func (c *CatRepoMock) FindAll() ([]entity.Cat, error) {
	args := c.Called()
	return args.Get(0).([]entity.Cat), args.Error(1)
}

func (c *CatRepoMock) DeleteById(id int) error {
	args := c.Called(id)
	return args.Error(0)
}