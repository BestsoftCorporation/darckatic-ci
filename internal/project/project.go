package project

import (
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name       string
	Server     server.RemoteServer     `gorm:"ForeignKey:ID"`
	Repository []repository.Repository `gorm:"many2many:project_repository;ForeignKey:ID"`
}
