package cfg

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// Load loads the configuration from a file into a struct.
func Load(file string, c any, opts ...viper.Option) error {
	var stat os.FileInfo

	stat, err := os.Stat(file)
	if err != nil {
		return err
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", file)
	}

	v := viper.NewWithOptions(opts...)
	v.SetConfigFile(file)

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(c)
	if err != nil {
		return err
	}

	return nil
}

// MustLoad loads the configuration from a file into a struct and panics on error.
func MustLoad(file string, c any, opts ...viper.Option) {
	if err := Load(file, c, opts...); err != nil {
		log.Fatalf("failed to load configuration: %s, error: %s", file, err)
	}
}
