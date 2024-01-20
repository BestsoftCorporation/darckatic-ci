package environment

import (
	"darkatic-ci/internal/project"
	"darkatic-ci/internal/server"
	"github.com/jinzhu/gorm"
)

type Environment struct {
	gorm.Model
	Name    string
	Server  server.RemoteServer `gorm:"foreignkey:ID"`
	Project []project.Project   `gorm:"many2many:environment_projects;ForeignKey:ID"`
	Parent  *Environment        `gorm:"foreignkey:EnvironmentID"`
}
