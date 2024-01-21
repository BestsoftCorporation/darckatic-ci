package repository

import (
	"darkatic-ci/internal/server"
	"darkatic-ci/internal/source"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	gorm.Model
	Name       string
	Branch     string
	RemotePath string
	Server     server.RemoteServer `gorm:"foreignkey:ID"`
	Source     source.Source       `gorm:"foreignkey:ID"`
	EnvVars    []EnvVars           `gorm:"many2many:repository_env_vars;foreignkey:ID"`
}

type EnvVars struct {
	gorm.Model
	Name  string
	Value string
}
