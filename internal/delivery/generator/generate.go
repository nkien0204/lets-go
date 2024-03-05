package generator

func (onl *delivery) HandleGenerate() error {
	return onl.gen.Generate()
}
