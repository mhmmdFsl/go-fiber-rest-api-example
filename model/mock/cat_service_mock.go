package mock

import (
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/stretchr/testify/mock"
)

type CatServiceMock struct {
	mock.Mock
}

func (c *CatServiceMock) AddNewCat(addCatDto *model.AddCatDto) (*entity.Cat, error) {
	args := c.Called(addCatDto)
	a0 := args.Get(0)
	a1 := args.Error(1)
	return a0.(*entity.Cat), a1
}
