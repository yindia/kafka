package github

// UserProfile represents a GitHub user's profile
type UserProfile struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	// Add more fields as needed
}

// OrgProfile represents a GitHub organization's profile
type OrgProfile struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	// Add more fields as needed
}

// Repository represents a GitHub repository
type Repository struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	FullName string      `json:"full_name"`
	Private  bool        `json:"private"`
	Owner    UserProfile `json:"owner"`
	// Add more fields as needed
}

// Commit represents a GitHub commit
type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
	// Add more fields as needed
}
