package provider

import (
	"context"
	"fmt"
	"github.com/google/go-github/v39/github"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
	"io"
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

func (g *GitHubProvider) ListBranches(owner, repo, token string) ([]string, error) {
	// Create a new GitHub client with OAuth2 authentication
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	result := []string{}

	// List branches
	branches, _, err := client.Repositories.ListBranches(ctx, owner, repo, nil)
	if err != nil {
		return result, fmt.Errorf("error listing branches: %v", err)
	}

	for _, branch := range branches {
		result = append(result, fmt.Sprintf("%s (SHA: %s)", *branch.Name, *branch.Commit.SHA))
	}

	return result, nil
}

// DownloadZip downloads a zip archive of a GitHub repository.
func (g *GitHubProvider) DownloadZip(owner, repo, token, destination string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	opts := &github.RepositoryContentGetOptions{
		"refs/heads/main",
	}

	zipURL, _, err := client.Repositories.GetArchiveLink(ctx, owner, repo, "zipball", opts, false)
	if err != nil {
		return err
	}

	url := "https://codeload.github.com/" + zipURL.Path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	req.Header.Set("Authorization", "token "+token)

	httpClient := http.Client{}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return nil
	}

	// Create the file
	out, err := os.Create("test.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)

	return nil
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

func (g *GitLabProvider) ListBranches(projectID interface{}) error {
	// Replace with your GitLab API token
	token := "YOUR_GITLAB_TOKEN"

	// Create a new GitLab client
	gitlabClient, err := gitlab.NewClient(token)
	if err != nil {
		return fmt.Errorf("error creating GitLab client: %v", err)
	}

	// Replace with your GitLab server URL
	//gitlabClient.set("https://gitlab.com/api/v4")

	// Get the branches of the project
	branches, _, err := gitlabClient.Branches.ListBranches(projectID, nil)
	if err != nil {
		return fmt.Errorf("error listing branches: %v", err)
	}

	// Print the branch names
	fmt.Println("Branches:")
	for _, branch := range branches {
		fmt.Printf("- %s (SHA: %s)\n", branch.Name, branch.Commit.ID)
	}

	return nil
}
