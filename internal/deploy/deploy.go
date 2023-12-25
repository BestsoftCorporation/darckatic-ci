package deploy

import (
	"darkatic-ci/internal/provider"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/source"
)

func Deploy(repo *repository.Repository) error {

	var prov provider.ProjectProvider

	src := repo.Source

	if src.SourceType == source.GitHub {
		prov = &provider.GitHubProvider{}
	} else if src.SourceType == source.GitLab {
		prov = &provider.GitLabProvider{}
	}

	err := prov.DownloadZip(src.Name, repo.Name, src.Token, repo.Name)
	if err != nil {
		return err
	}

	err = repo.Server.CopyFileToRemote(repo.Name+".zip", repo.RemotePath+"/"+repo.Name+".zip")
	if err != nil {
		return err
	}

	err = repo.Server.UnzipFileOnRemote(repo.RemotePath+"/"+repo.Name+".zip", repo.RemotePath+"/"+repo.Name)
	if err != nil {
		return err
	}

	return nil
}
