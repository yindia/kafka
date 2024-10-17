package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.github.com"

// Client struct to hold the HTTP client
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new GitHub API client
func NewClient() *Client {
	return &Client{httpClient: &http.Client{}}
}

// UserProfile retrieves a user's profile
func (c *Client) UserProfile(username string) (*UserProfile, error) {
	url := fmt.Sprintf("%s/users/%s", baseURL, username)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var profile UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

// OrgProfile retrieves an organization's profile
func (c *Client) OrgProfile(org string) (*OrgProfile, error) {
	url := fmt.Sprintf("%s/orgs/%s", baseURL, org)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var organization OrgProfile
	if err := json.NewDecoder(resp.Body).Decode(&organization); err != nil {
		return nil, err
	}
	return &organization, nil
}

// ListRepositories lists repositories for a user or organization with pagination
func (c *Client) ListRepositories(owner string, page int) ([]Repository, error) {
	url := fmt.Sprintf("%s/users/%s/repos?page=%d", baseURL, owner, page)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}

// GetRepository retrieves a specific repository
func (c *Client) GetRepository(owner, repo string) (*Repository, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", baseURL, owner, repo)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repository Repository
	if err := json.NewDecoder(resp.Body).Decode(&repository); err != nil {
		return nil, err
	}
	return &repository, nil
}

// GetCommits retrieves commits for a repository with pagination
func (c *Client) GetCommits(owner, repo string, page int) ([]Commit, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/commits?page=%d", baseURL, owner, repo, page)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}
	return commits, nil
}
