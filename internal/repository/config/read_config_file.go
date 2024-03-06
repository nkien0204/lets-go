package config

import (
	"fmt"
	"os"

	"github.com/nkien0204/lets-go/internal/domain/entity/config"
	"gopkg.in/yaml.v2"
)

func (repo *repository) ReadConfigFile() (result config.ConfigFileReadResponseEntity, err error) {
	buf, err := os.ReadFile(repo.fileName)
	if err != nil {
		return result, err
	}

	c := &config.Cfg{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return result, fmt.Errorf("in file %q: %w", repo.fileName, err)
	}
	result.Config = c

	return result, nil
}
