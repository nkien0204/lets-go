package greeting_test

import (
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
