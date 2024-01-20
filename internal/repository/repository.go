package repository

import (
	"darkatic-ci/internal/server"
	"darkatic-ci/internal/source"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	gorm.Model
	SourceID   uint
	ServerID   uint
	Name       string
	Branch     string
	RemotePath string
	Server     server.RemoteServer `gorm:"foreignkey:ServerID"`
	Source     source.Source       `gorm:"foreignkey:SourceID"`
	EnvVars    []EnvVars           `gorm:"foreignkey:EnvVarID"`
}

type EnvVars struct {
	gorm.Model
	Name  string
	Value string
}
