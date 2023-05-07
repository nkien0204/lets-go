package confi

import (
	"fmt"
	"os"

	"github.com/nkien0204/lets-go/internal/entity/config"
	"gopkg.in/yaml.v2"
)

type usecase struct {
	cfg *config.Cfg
}

func NewConfig() *usecase {
	return &usecase{}
}

func (u *usecase) LoadConfig() *config.Cfg {
	var err error
	if u.cfg, err = newConfig(); err != nil {
		panic(err)
	}
	return u.cfg
}

func newConfig() (*config.Cfg, error) {
	return readConf(config.CONFIG_FILENAME)
}

func readConf(filename string) (*config.Cfg, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &config.Cfg{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}
