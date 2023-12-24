package main

import (
	"darkatic-ci/cmd"

	"darkatic-ci/internal/api"
	"darkatic-ci/internal/deploy"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
	"darkatic-ci/internal/source"
	"fmt"
)

func main() {

	ser := &server.RemoteServer{
		Host:       "104.156.225.68",
		Port:       "22",
		Username:   "root",
		Key:        "2Lw+3@}T1tnk}5Nh",
		AuthMethod: server.Password,
	}

	own := &source.Source{
		Token: "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",
		Name:  "BestsoftCorporation",
	}

	repo := &repository.Repository{
		Name:       "Darkatic-Website",
		RemotePath: "/root",
		Source:     *own,
	}

	err := deploy.Deploy(ser, repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	api.StartServer()
	cmd.Execute()
}
