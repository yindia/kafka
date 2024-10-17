package github

import (
	"context"
	"log"

	"github.com/google/go-github/v39/github" // Import the Google GitHub package
	"golang.org/x/oauth2"                    // Import OAuth2 for authentication
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)                  // Include date, time, and file info in logs
	log.Println("Logging initialized: All logs will be printed.") // Log initialization message
}

const baseURL = "https://api.github.com"

// Client struct to hold the GitHub client
type Client struct {
	githubClient *github.Client
}

// NewClient creates a new GitHub API client
func NewClient(token string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &Client{githubClient: github.NewClient(tc)}
}

// UserProfile retrieves a user's profile
func (c *Client) UserProfile(username string) (*github.User, error) {
	user, _, err := c.githubClient.Users.Get(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// OrgProfile retrieves an organization's profile
func (c *Client) OrgProfile(org string) (*github.Organization, error) {
	orgProfile, _, err := c.githubClient.Organizations.Get(context.Background(), org)
	if err != nil {
		return nil, err
	}
	return orgProfile, nil
}

// ListRepositories lists repositories for a user or organization with pagination
func (c *Client) ListRepositories(owner string, page, perPage int) ([]*github.Repository, error) {
	repos, _, err := c.githubClient.Repositories.List(context.Background(), owner, &github.RepositoryListOptions{
		ListOptions: github.ListOptions{Page: page, PerPage: perPage},
	})
	if err != nil {
		return nil, err
	}
	return repos, nil
}

// GetRepository retrieves a specific repository
func (c *Client) GetRepository(owner, repo string) (*github.Repository, error) {
	repository, _, err := c.githubClient.Repositories.Get(context.Background(), owner, repo)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

// GetCommits retrieves commits for a repository with pagination
func (c *Client) GetCommits(owner, repo string, page int) ([]*github.RepositoryCommit, error) {
	commits, _, err := c.githubClient.Repositories.ListCommits(context.Background(), owner, repo, &github.CommitsListOptions{
		ListOptions: github.ListOptions{Page: page},
	})
	if err != nil {
		return nil, err
	}
	return commits, nil
}
