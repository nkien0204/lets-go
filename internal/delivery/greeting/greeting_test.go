package greeting_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	delivery "github.com/nkien0204/lets-go/internal/delivery/greeting"
	"github.com/nkien0204/lets-go/internal/domain/entity/greeting"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/stretchr/testify/assert"
)

func TestGreetingHappy(t *testing.T) {
	expectMsg := "hello, world!"
	usecase := mock.NewGreetingUsecase(t)
	usecase.On("Greeting").Return(greeting.GreetingResponseEntity{Msg: expectMsg}, nil)

	w := httptest.ResponseRecorder{}

	delivery := delivery.NewDelivery(usecase)
	handler := delivery.Greeting()
	handler(&w, nil)
	assert.NotNil(t, w)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGreetingError(t *testing.T) {
	usecase := mock.NewGreetingUsecase(t)
	usecase.On("Greeting").Return(greeting.GreetingResponseEntity{}, errors.New("service error"))

	w := httptest.ResponseRecorder{}

	delivery := delivery.NewDelivery(usecase)
	handler := delivery.Greeting()
	handler(&w, nil)

	// The function sets status 500 but continues to execute, causing the final status to be determined by ResponseWriter behavior
	// Let's check what actually happens
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	// The response body should contain some content
	bodyStr := w.Body.String()
	assert.NotEmpty(t, bodyStr)
}
