package project

import (
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"github.com/jinzhu/gorm"
)

type Project struct {
	gorm.Model
	Name       string
	Server     server.RemoteServer
	Repository []repository.Repository
}
