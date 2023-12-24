package darkatic_ci

import (
	"darkatic-ci/internal/deploy"
	"darkatic-ci/internal/owner"
	"darkatic-ci/internal/provider"
	"darkatic-ci/internal/repository"
	"darkatic-ci/internal/server"
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

	prov := &provider.GitHubProvider{}
	own := &owner.Owner{
		Token: "",
		Name:  "BestsoftCorporation",
	}

	repo := &repository.Repository{
		Name:       "Darkatic-Website",
		RemotePath: "/root",
		Provider:   prov,
	}

	err := deploy.Deploy(ser, own, repo)
	if err != nil {
		fmt.Println(err)
		return
	}
}
