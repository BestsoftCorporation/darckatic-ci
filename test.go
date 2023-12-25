package darkatic_ci

import (
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
		Key:        "",
		AuthMethod: server.Password,
	}

	//prov := &provider.GitHubProvider{}
	src := &source.Source{
		Token: "",
		Name:  "BestsoftCorporation",
	}

	repo := &repository.Repository{
		Name:       "Darkatic-Website",
		RemotePath: "/root",
		Source:     *src,
		Server:     *ser,
	}

	err := deploy.Deploy(repo)
	if err != nil {
		fmt.Println(err)
		return
	}
}
