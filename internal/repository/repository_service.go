package repository

import (
	"darkatic-ci/internal/db"
	"fmt"
)

func init() {
	db.DB.AutoMigrate(&Repository{})
}

func GetRepositoryByName(name string) (*Repository, error) {
	var repository Repository
	if err := db.DB.Where("Name = ?", name).First(&repository).Error; err != nil {
		return nil, fmt.Errorf("repository not found")
	}

	return &repository, nil
}

func GetRepositoryById(id string) (*Repository, error) {
	var repository Repository
	if err := db.DB.Preload("Source").First(&repository, id).Error; err != nil {
		return nil, fmt.Errorf("repository not found")
	}

	return &repository, nil
}
