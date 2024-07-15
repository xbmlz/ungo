package db

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Driver string `json:"driver" yaml:"driver" env:"DB_DRIVER" default:"sqlite"`
	Host   string `json:"host" yaml:"host" env:"DB_HOST" default:""`
	Port   string `json:"port" yaml:"port" env:"DB_PORT" default:""`
	User   string `json:"user" yaml:"user" env:"DB_USER" default:""`
	Pass   string `json:"pass" yaml:"pass" env:"DB_PASS" default:""`
	Name   string `json:"name" yaml:"name" env:"DB_NAME" default:"./db.sqlite"`
}

func (c *Config) DSN() string {
	if c.Driver == "sqlite" {
		return c.Name
	}
	return c.Driver + "://" + c.User + ":" + c.Pass + "@" + c.Host + ":" + c.Port + "/" + c.Name
}

func Connect(config Config, opts ...gorm.Option) (db *gorm.DB, err error) {

	if config.Driver == "sqlite" {
		// if config.Name path is not exist, create it
		if _, err := os.Stat(config.Name); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(config.Name), os.ModePerm)
		}
	}

	switch config.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Name))
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.DSN()), &gorm.Config{})
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
	default:
		err = errors.New("unsupported driver")
	}

	return db, err
}

func MustConnect(config Config, opts ...gorm.Option) *gorm.DB {
	db, err := Connect(config, opts...)
	if err != nil {
		panic(err)
	}
	return db
}
