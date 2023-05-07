package generator_test

import (
	"testing"

	delivery "github.com/nkien0204/lets-go/internal/delivery/generator"
	"github.com/nkien0204/lets-go/internal/delivery/generator/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	gen := mocks.NewGenerateBehaviors(t)
	gen.On("Generate").Return(nil)

	genDelivery := delivery.NewDelivery(gen)
	err := genDelivery.HandleGenerate()

	assert.Nil(t, err)
}
