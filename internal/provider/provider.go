package provider

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// ProjectProvider defines the interface for fetching projects and downloading zip archives from different providers.
type ProjectProvider interface {
	FetchProjects(owner, token string) ([]string, error)
	DownloadZip(owner, repo, token, destination string) error
}

// GitHubProvider is a GitHub API implementation of the ProjectProvider interface.
type GitHubProvider struct{}

// FetchProjects fetches projects from GitHub API.
func (g *GitHubProvider) FetchProjects(owner, token string) ([]string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	repos, _, err := client.Repositories.List(ctx, owner, nil)
	if err != nil {
		return nil, err
	}

	var projectNames []string
	for _, repo := range repos {
		projectNames = append(projectNames, *repo.Name)
	}

	return projectNames, nil
}

// DownloadZip downloads a zip archive of a GitHub repository.
func (g *GitHubProvider) DownloadZip(owner, repo, token, destination string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	zipURL, _, err := client.Repositories.GetArchiveLink(ctx, owner, repo, github.Zipball, nil, true)
	if err != nil {
		return err
	}

	resp, err := http.Get(zipURL.Path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	zipData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	zipPath := filepath.Join(destination, fmt.Sprintf("%s-%s.zip", owner, repo))
	return ioutil.WriteFile(zipPath, zipData, os.ModePerm)
}

// GitLabProvider is a GitLab API implementation of the ProjectProvider interface.
type GitLabProvider struct{}

// FetchProjects fetches projects from GitLab API.
func (g *GitLabProvider) FetchProjects(owner, token string) ([]string, error) {
	client, err := gitlab.NewClient(token)
	if err != nil {
		return nil, err
	}

	projects, _, err := client.Groups.ListGroupProjects(owner, nil)
	if err != nil {
		return nil, err
	}

	var projectNames []string
	for _, project := range projects {
		projectNames = append(projectNames, project.Name)
	}

	return projectNames, nil
}

// DownloadZip downloads a zip archive of a GitLab repository.
func (g *GitLabProvider) DownloadZip(owner, repo, token, destination string) error {
	client, err := gitlab.NewClient(token)
	if err != nil {
		return err
	}

	opt := &gitlab.ArchiveOptions{
		// Replace with the desired branch or commit reference
	}

	content, _, err := client.Repositories.Archive(owner, opt, nil)
	if err != nil {
		return err
	}

	zipPath := filepath.Join(destination, fmt.Sprintf("%s-%s.zip", owner, repo))
	return ioutil.WriteFile(zipPath, content, os.ModePerm)
}
