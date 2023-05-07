package generators

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/usecases/generators/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
    gen := mocks.NewGenerateBehaviors(t)
    gen.On("Generate").Return(nil)

    genUseCase := NewGenUseCase(gen)
    err := genUseCase.HandleGenerate()

    assert.Nil(t, err)
}
