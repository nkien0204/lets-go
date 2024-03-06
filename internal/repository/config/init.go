package config

type repository struct {
	fileName string
}

func NewRepository(fileName string) *repository {
	return &repository{
		fileName: fileName,
	}
}
