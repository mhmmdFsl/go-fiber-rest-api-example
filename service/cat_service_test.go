package service

import (
	"testing"
	"time"

	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/httperror"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/mock"
	"github.com/stretchr/testify/assert"
)

const (
	Save     string = "Save"
	FindById string = "FindById"
	FindAll  string = "FindAll"
)

func TestAddCatService(t *testing.T) {

	tests := []struct {
		description          string
		param                interface{}
		mockMethodName       string
		mockParams           *entity.Cat
		mockExpectedResponse *entity.Cat
		mockExpectedError    error
		expectedError        bool
	}{
		{
			description: "add cat",
			param: model.AddCatDto{
				Name:  "Test cat",
				Breed: "Test breed",
				Color: "Black",
			},
			mockMethodName: Save,
			mockParams: &entity.Cat{
				ID:        0,
				Name:      "Test cat",
				Breed:     "Test breed",
				Color:     "Black",
				BirthAt:   time.Time{},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			mockExpectedResponse: &entity.Cat{
				ID:        1,
				Name:      "Test cat",
				Breed:     "Test breed",
				Color:     "Black",
				BirthAt:   time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		{
			description: "add cat should error",
			param: model.AddCatDto{
				Name:  "Test cat",
				Breed: "Test breed",
				Color: "Orange",
			},
			mockMethodName: Save,
			mockParams: &entity.Cat{
				ID:        0,
				Name:      "Test cat",
				Breed:     "Test breed",
				Color:     "Orange",
				BirthAt:   time.Time{},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			mockExpectedError: &httperror.HttpError{
				Message:   "failed add new cat",
				Status:    "FAILED",
				ErrorType: httperror.BAD_REQUEST,
			},
			expectedError: true,
		},
	}

	repo := &mock.CatRepoMock{}
	service := NewCatService(&CatServiceCfg{
		CatRepo: repo,
	})

	for _, test := range tests {

		var cat *entity.Cat
		var err error
		repo.On(test.mockMethodName, test.mockParams).Return(test.mockExpectedResponse, test.mockExpectedError)
		param := test.param.(model.AddCatDto)
		cat, err = service.AddNewCat(&param)

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.mockExpectedResponse, cat, test.description)
	}
}

func TestGetCatById(t *testing.T) {
	tests := []struct {
		description          string
		param                int
		mockMethodName       string
		mockParams           int
		mockExpectedResponse *entity.Cat
		mockExpectedError    error
		expectedError        bool
	}{
		{
			description:    "get cat by id",
			param:          1,
			mockMethodName: FindById,
			mockParams:     1,
			mockExpectedResponse: &entity.Cat{
				ID:        1,
				Name:      "Test cat",
				Breed:     "Test breed",
				Color:     "Black",
				BirthAt:   time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: false,
		},
		{
			description:    "get cat by id not found",
			param:          2,
			mockMethodName: FindById,
			mockParams:     2,
			mockExpectedError: &httperror.HttpError{
				Message:   "cat with id 2 not found",
				Status:    "FAILED",
				ErrorType: httperror.NOT_FOUND,
			},
			expectedError: true,
		},
	}

	repo := &mock.CatRepoMock{}
	service := NewCatService(&CatServiceCfg{
		CatRepo: repo,
	})

	for _, test := range tests {

		var cat *entity.Cat
		var err error
		repo.On(test.mockMethodName, test.mockParams).Return(test.mockExpectedResponse, test.mockExpectedError)
		cat, err = service.GetById(test.param)

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.mockExpectedResponse, cat, test.description)
	}
}

func TestGetAllCat(t *testing.T) {
	tests := []struct {
		description          string
		mockMethodName       string
		mockExpectedResponse []entity.Cat
		mockExpectedError    error
		expectedError        bool
	}{
		{
			description:    "get all cat",
			mockMethodName: FindAll,
			mockExpectedResponse: []entity.Cat{
				{
					ID:        1,
					Name:      "Test Cat",
					Breed:     "Kampung",
					Color:     "Orange",
					BirthAt:   time.Now(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			expectedError: false,
		},
	}

	repo := &mock.CatRepoMock{}
	service := NewCatService(&CatServiceCfg{
		CatRepo: repo,
	})

	for _, test := range tests {

		var rs []entity.Cat
		var err error
		repo.On(test.mockMethodName).Return(test.mockExpectedResponse, test.mockExpectedError)
		rs, err = service.GetAll()
		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.mockExpectedResponse, rs, test.description)
	}
}
