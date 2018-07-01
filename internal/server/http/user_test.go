package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fhbkoller/workshop-go/domain"
	"github.com/stretchr/testify/assert"
)

type UserServiceMock struct {
	RetrieveFunc func(userID string) (*domain.User, error)
	CreateFunc   func(user *domain.User) error
}

func (usm *UserServiceMock) Retrieve(userID string) (*domain.User, error) {
	return usm.RetrieveFunc(userID)
}

func (usm *UserServiceMock) Create(user *domain.User) error {
	return usm.CreateFunc(user)
}

func TestController(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		userServiceMock := &UserServiceMock{
			RetrieveFunc: func(userID string) (*domain.User, error) {
				return domain.NewUser("1", "Fernando", 0, nil, nil), nil
			},
		}

		router := NewHandler(userServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/users/1"

		req, _ := http.NewRequest("GET", endpoint, nil)

		router.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code)

		body, err := ioutil.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		expectedBody := []byte(`{"id": "1", "name": "Fernando"}`)

		assert.JSONEq(t, string(expectedBody), string(body))

	})
	t.Run("fail", func(t *testing.T) {
		userServiceMock := &UserServiceMock{
			RetrieveFunc: func(userID string) (*domain.User, error) {
				return nil, errors.New("error")
			},
		}

		router := NewHandler(userServiceMock)

		response := httptest.NewRecorder()
		endpoint := "/v1/users/1"

		req, _ := http.NewRequest("GET", endpoint, nil)

		router.ServeHTTP(response, req)

		assert.Equal(t, http.StatusInternalServerError, response.Code)

		body, err := ioutil.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.Empty(t, body)
	})
}
