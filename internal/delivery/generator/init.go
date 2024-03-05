package generator

type GeneratorBehaviors interface {
	Generate() error
}

type delivery struct {
	gen GeneratorBehaviors
}

func NewDelivery(o GeneratorBehaviors) *delivery {
	return &delivery{gen: o}
}
