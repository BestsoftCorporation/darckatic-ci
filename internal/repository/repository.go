package repository

import "darkatic-ci/internal/provider"

type Repository struct {
	Name       string
	RemotePath string
	Provider   provider.ProjectProvider
}
