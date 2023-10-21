package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/entity"
	"github.com/mhmmdFsl/go-fiber-rest-api-example/model/mock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHandlerCat(t *testing.T) {

	tests := []struct {
		description      string
		route            string // input route
		expectedError    bool
		expectedCode     int
		httpMethod       string
		requestBody      *model.AddCatDto
		expectedResponse *entity.Cat
	}{
		{
			description:   "Hello from cat",
			route:         "/api/v1/cat/hello",
			expectedError: false,
			expectedCode:  200,
			httpMethod:    "GET",
		},
		{
			description:   "Invalid uri",
			route:         "/api/v1/dog/hello",
			expectedError: false,
			expectedCode:  404,
			httpMethod:    "GET",
		},
		{
			description:   "add new cat",
			route:         "/api/v1/cat/add",
			expectedError: false,
			expectedCode:  200,
			httpMethod:    "POST",
			requestBody: &model.AddCatDto{
				Name:  "Test cat",
				Breed: "Test breed",
				Color: "Black",
			},
			expectedResponse: &entity.Cat{
				ID:        1,
				Name:      "Test cat",
				Breed:     "Test breed",
				Color:     "Black",
				BirthAt:   time.Now(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	app := fiber.New()

	mockService := mock.CatServiceMock{}
	cfg := &HandlerCfg{
		App:        app,
		CatService: &mockService,
	}

	NewHandler(cfg)

	for _, test := range tests {

		// setup mock
		if test.requestBody != nil {
			mockService.On("AddNewCat", test.requestBody).Return(test.expectedResponse, nil)

		}
		j, _ := json.Marshal(test.requestBody)
		reader := strings.NewReader(string(j))

		req := httptest.NewRequest(test.httpMethod, test.route, reader)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
