package source

import "github.com/jinzhu/gorm"

type SourceType int

const (
	GitHub SourceType = iota
	GitLab
)

type Source struct {
	gorm.Model
	Token      string
	Name       string
	SourceType SourceType
}
