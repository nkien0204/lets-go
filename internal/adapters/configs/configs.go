package configs

import (
	"fmt"
	"os"

	"github.com/nkien0204/lets-go/internal/entities/configs"
	"gopkg.in/yaml.v2"
)

type configAdapter struct {
	cfg *configs.Cfg
}

func NewConfigs() *configAdapter {
	return &configAdapter{}
}

func (c *configAdapter) LoadConfigs() *configs.Cfg {
	var err error
	if c.cfg, err = newConfigs(); err != nil {
		panic(err)
	}
	return c.cfg
}

func newConfigs() (*configs.Cfg, error) {
	return readConf(configs.CONFIG_FILENAME)
}

func readConf(filename string) (*configs.Cfg, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &configs.Cfg{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}
