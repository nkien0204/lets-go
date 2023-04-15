package drivers

import (
	"fmt"
	"io/ioutil"

	"github.com/nkien0204/lets-go/internal/entities"
	"gopkg.in/yaml.v2"
)

type configs struct {
	cfg *entities.Cfg
}

func NewConfigs() *configs {
	return &configs{}
}

func (c *configs) LoadConfigs() *entities.Cfg {
	var err error
	if c.cfg, err = newConfigs(); err != nil {
		panic(err)
	}
	return c.cfg
}

func newConfigs() (*entities.Cfg, error) {
	return readConf(entities.CONFIG_FILENAME)
}

func readConf(filename string) (*entities.Cfg, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &entities.Cfg{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}
