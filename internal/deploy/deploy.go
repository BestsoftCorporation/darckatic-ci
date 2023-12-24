package deploy

import (
	"darkatic-ci/internal/owner"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
)

func Deploy(server *server.RemoteServer, owner *owner.Owner, repo *repository.Repository) error {

	err := repo.Provider.DownloadZip(owner.Name, repo.Name, owner.Token, repo.Name)
	if err != nil {
		return err
	}

	err = server.CopyFileToRemote(repo.Name+".zip", repo.RemotePath+"/"+repo.Name+".zip")
	if err != nil {
		return err
	}

	err = server.UnzipFileOnRemote(repo.RemotePath+"/"+repo.Name+".zip", repo.RemotePath+"/"+repo.Name)
	if err != nil {
		return err
	}

	return nil
}
