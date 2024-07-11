package unconf

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config interface {
	Parse(any) error
	Get(string) any
	GetBool(string) bool
	GetString(string) string
	GetInt(string) int
}

type Conf struct {
	Path   string `json:"path"`
	parser *viper.Viper
}

func New(filePath string) (c Config, err error) {
	var stat os.FileInfo

	stat, err = os.Stat(filePath)
	if err != nil {
		return
	}

	if !stat.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", filePath)
	}

	p := viper.New()
	p.SetConfigFile(filePath)

	err = p.ReadInConfig()
	if err != nil {
		return
	}

	return &Conf{Path: filePath, parser: p}, nil
}

// Parse parses the configuration by object pointer
func (c *Conf) Parse(obj any) error {
	return c.parser.Unmarshal(obj)
}

func (c *Conf) GetString(key string) string {
	return c.parser.GetString(key)
}

func (c *Conf) GetBool(key string) bool {
	return c.parser.GetBool(key)
}

func (c *Conf) Get(s string) any {
	return c.parser.Get(s)
}

func (c *Conf) GetInt(s string) int {
	return c.parser.GetInt(s)
}
