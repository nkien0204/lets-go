package generators

type GenerateBehaviors interface {
    Generate() error 
}

type generateUseCase struct {
    gen GenerateBehaviors
}

func NewGenUseCase(o GenerateBehaviors) *generateUseCase {
    return &generateUseCase{gen: o}
}

func (onl *generateUseCase) HandleGenerate() error {
    return onl.gen.Generate()
}
