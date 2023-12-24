package repository

import (
	"darkatic-ci/internal/source"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	gorm.Model
	SourceID   uint
	Name       string
	Branch     string
	RemotePath string
	Source     source.Source
}
