package greeting_test

import (
	"errors"
	"testing"

	greetingEntity "github.com/nkien0204/lets-go/internal/domain/entity/greeting"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/nkien0204/lets-go/internal/usecase/greeting"
	"github.com/stretchr/testify/assert"
)

func TestGreetingHappy(t *testing.T) {
	repo := mock.NewGreetingRepository(t)
	repo.On(
		"SayHello",
	).Return(greetingEntity.GreetingResponseEntity{}, nil)

	usecase := greeting.NewUsecase(repo)
	_, err := usecase.Greeting()

	assert.Nil(t, err)
}

func TestGreetingError(t *testing.T) {
	repo := mock.NewGreetingRepository(t)
	repo.On(
		"SayHello",
	).Return(greetingEntity.GreetingResponseEntity{}, errors.New("something went wrong"))

	usecase := greeting.NewUsecase(repo)
	_, err := usecase.Greeting()

	assert.Error(t, err)
}
