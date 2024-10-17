package cmd

import (
	"fmt"
	"strings"

	v1 "kafka/internal/gen/cloud/v1"
	"kafka/internal/github"
	"kafka/internal/kafka" // Import Kafka package

	gh "github.com/google/go-github/v39/github"
	"github.com/spf13/cobra"
)

var (
	profileType       string          // "user" or "org"
	usernames         string          // Comma-separated usernames
	producer          *kafka.Producer // Kafka producer
	brokerURL         string          // Kafka broker URL
	schemaRegistryURL string          // Schema registry URL
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Run the workflow orchestration server",
	Long:    `Start the workflow orchestration server and continuously stream task updates.`,
	Example: `  task fetch --type user --usernames user1,user2 --broker-url <broker-url>`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := github.NewClient("")

		for _, username := range []string{"yindia"} {
			if profileType == "user" {
				profile, err := client.UserProfile(strings.TrimSpace(username))
				if err != nil {
					return err
				}
				fmt.Printf("User Profile: %+v\n", profile.ID)
				err = producer.SendProfile("profiles", &v1.Profile{
					Profile: &v1.Profile_User{
						User: &v1.User{
							Login:                   profile.GetLogin(),
							Id:                      int64(profile.GetID()),
							AvatarUrl:               profile.GetAvatarURL(),
							HtmlUrl:                 profile.GetHTMLURL(),
							Type:                    "user",
							Email:                   profile.GetEmail(),                                                                                                // Add email field
							Bio:                     profile.GetBio(),                                                                                                  // Add bio field
							Blog:                    profile.GetBlog(),                                                                                                 // Add blog field
							PublicRepos:             int32(profile.GetPublicRepos()),                                                                                   // Add public repos count
							PublicGists:             int32(profile.GetPublicGists()),                                                                                   // Add public gists count
							Followers:               int32(profile.GetFollowers()),                                                                                     // Add followers count
							Following:               int32(profile.GetFollowing()),                                                                                     // Add following count
							GravatarId:              profile.GetGravatarID(),                                                                                           // Optional, fetched from GitHub
							Name:                    profile.GetName(),                                                                                                 // Optional, fetched from GitHub
							Company:                 profile.GetCompany(),                                                                                              // Optional, fetched from GitHub
							Location:                profile.GetLocation(),                                                                                             // Optional, fetched from GitHub
							Hireable:                profile.GetHireable(),                                                                                             // Optional, fetched from GitHub
							TwitterUsername:         profile.GetTwitterUsername(),                                                                                      // Optional, fetched from GitHub
							TotalPrivateRepos:       int32(profile.GetTotalPrivateRepos()),                                                                             // Optional, fetched from GitHub
							OwnedPrivateRepos:       int32(profile.GetOwnedPrivateRepos()),                                                                             // Optional, fetched from GitHub
							PrivateGists:            int32(profile.GetPrivateGists()),                                                                                  // Optional, fetched from GitHub
							DiskUsage:               int32(profile.GetDiskUsage()),                                                                                     // Optional, fetched from GitHub
							Collaborators:           int32(profile.GetCollaborators()),                                                                                 // Optional, fetched from GitHub
							TwoFactorAuthentication: profile.GetTwoFactorAuthentication(),                                                                              // Optional, fetched from GitHub
							Url:                     profile.GetURL(),                                                                                                  // Optional, fetched from GitHub
							EventsUrl:               profile.GetEventsURL(),                                                                                            // Optional, fetched from GitHub
							FollowingUrl:            profile.GetFollowingURL(),                                                                                         // Optional, fetched from GitHub
							FollowersUrl:            profile.GetFollowersURL(),                                                                                         // Optional, fetched from GitHub
							GistsUrl:                profile.GetGistsURL(),                                                                                             // Optional, fetched from GitHub
							OrganizationsUrl:        profile.GetOrganizationsURL(),                                                                                     // Optional, fetched from GitHub
							ReceivedEventsUrl:       profile.GetReceivedEventsURL(),                                                                                    // Optional, fetched from GitHub
							ReposUrl:                profile.GetReposURL(),                                                                                             // Optional, fetched from GitHub
							StarredUrl:              profile.GetStarredURL(),                                                                                           // Optional, fetched from GitHub
							SubscriptionsUrl:        profile.GetSubscriptionsURL(),                                                                                     // Optional, fetched from GitHub
							Permissions:             profile.GetPermissions(),                                                                                          // Optional, no data available                                                                                                            // Optional, no data available
							CreatedAt:               &v1.Timestamp{Seconds: int64(profile.GetCreatedAt().Second()), Nanos: int32(profile.GetCreatedAt().Nanosecond())}, // Convert to *cloudv1.Timestamp
							UpdatedAt:               &v1.Timestamp{Seconds: int64(profile.GetUpdatedAt().Second()), Nanos: int32(profile.GetUpdatedAt().Nanosecond())}, // Convert to *cloudv1.Timestamp
						},
					},
				}) // Send org profile to Kafka topic
				if err != nil {
					fmt.Printf("Failed to send message: %s\n", err)
					return err
				}
			} else if profileType == "org" {
				profile, err := client.OrgProfile(strings.TrimSpace(username))
				if err != nil {
					return err
				}

				err = producer.SendProfile("profiles", &v1.Profile{
					Profile: &v1.Profile_Organization{
						Organization: &v1.Organization{
							Login:            profile.GetLogin(),
							Id:               int64(profile.GetID()),
							Url:              profile.GetURL(),
							ReposUrl:         profile.GetReposURL(),
							EventsUrl:        profile.GetEventsURL(),
							HooksUrl:         profile.GetHooksURL(),
							IssuesUrl:        profile.GetIssuesURL(),
							MembersUrl:       profile.GetMembersURL(),
							PublicMembersUrl: profile.GetPublicMembersURL(),
							AvatarUrl:        profile.GetAvatarURL(),
							Description:      profile.GetDescription(),
							Name:             profile.GetName(),
							Company:          profile.GetCompany(),
							Blog:             profile.GetBlog(),
							Location:         profile.GetLocation(),
							Email:            profile.GetEmail(),
							TwitterUsername:  profile.GetTwitterUsername(),
							IsVerified:       profile.GetIsVerified(),
							CreatedAt:        &v1.Timestamp{Seconds: int64(profile.GetCreatedAt().Second()), Nanos: int32(profile.GetCreatedAt().Nanosecond())},
							UpdatedAt:        &v1.Timestamp{Seconds: int64(profile.GetUpdatedAt().Second()), Nanos: int32(profile.GetUpdatedAt().Nanosecond())},
						},
					},
				}) // Send org profile to Kafka topic
				if err != nil {
					fmt.Printf("Failed to send message: %s\n", err)
					return err
				}
			}

			page := 1
			perPage := 100
			var allRepos []*gh.Repository

			for {
				repos, err := client.ListRepositories(strings.TrimSpace(username), page, perPage) // Fetch repositories with pagination
				if err != nil {
					return err
				}
				if len(repos) == 0 {
					break // Exit loop if no more repositories are returned
				}
				allRepos = append(allRepos, repos...) // Append fetched repositories to the allRepos slice
				page++                                // Increment page number for the next request
			}

			for _, repo := range allRepos {

				data := &v1.Repository{
					Id:     int64(repo.GetID()), // Unique identifier for the repository
					NodeId: repo.GetNodeID(),    // Unique node identifier for the repository
					Owner: &v1.User{ // Owner of the repository
						Login: repo.GetOwner().GetLogin(),
						Id:    int64(repo.GetOwner().GetID()),
					},
					Name:                repo.GetName(),                                                                                              // Name of the repository
					FullName:            repo.GetFullName(),                                                                                          // Full name of the repository
					Description:         repo.GetDescription(),                                                                                       // Brief description of the repository
					Homepage:            repo.GetHomepage(),                                                                                          // Homepage URL of the repository
					DefaultBranch:       repo.GetDefaultBranch(),                                                                                     // Default branch of the repository
					CreatedAt:           &v1.Timestamp{Seconds: int64(repo.GetCreatedAt().Second()), Nanos: int32(repo.GetCreatedAt().Nanosecond())}, // Created timestamp
					PushedAt:            &v1.Timestamp{Seconds: int64(repo.GetPushedAt().Second()), Nanos: int32(repo.GetPushedAt().Nanosecond())},   // Pushed timestamp
					UpdatedAt:           &v1.Timestamp{Seconds: int64(repo.GetUpdatedAt().Second()), Nanos: int32(repo.GetUpdatedAt().Nanosecond())}, // Updated timestamp
					HtmlUrl:             repo.GetHTMLURL(),                                                                                           // HTML URL of the repository
					CloneUrl:            repo.GetCloneURL(),                                                                                          // Clone URL of the repository
					GitUrl:              repo.GetGitURL(),                                                                                            // Git URL of the repository
					MirrorUrl:           repo.GetMirrorURL(),                                                                                         // Mirror URL of the repository
					SshUrl:              repo.GetSSHURL(),                                                                                            // SSH URL of the repository
					SvnUrl:              repo.GetSVNURL(),                                                                                            // SVN URL of the repository
					Language:            repo.GetLanguage(),                                                                                          // Language for the repository
					Fork:                repo.GetFork(),                                                                                              // Number of forks for the repository
					ForksCount:          int32(repo.GetForksCount()),                                                                                 // Number of forks for the repository
					NetworkCount:        int32(repo.GetNetworkCount()),                                                                               // Number of network forks
					OpenIssuesCount:     int32(repo.GetOpenIssuesCount()),                                                                            // Number of open issues
					OpenIssues:          int32(repo.GetOpenIssues()),                                                                                 // Number of open issues for the repository
					StargazersCount:     int32(repo.GetStargazersCount()),                                                                            // Number of stargazers
					SubscribersCount:    int32(repo.GetSubscribersCount()),                                                                           // Number of subscribers
					Watchers:            int32(repo.GetWatchers()),                                                                                   // Number of watchers for the repository
					Size:                int32(repo.GetSize()),                                                                                       // Size of the repository
					HasProjects:         repo.GetHasProjects(),                                                                                       // Whether the repository has projects
					HasDownloads:        repo.GetHasDownloads(),                                                                                      // Whether the repository has downloads
					Archived:            repo.GetArchived(),                                                                                          // Whether the repository is archived
					Disabled:            repo.GetDisabled(),                                                                                          // Whether the repository is disabled
					AllowMergeCommit:    repo.GetAllowMergeCommit(),                                                                                  // Whether merge commits are allowed
					AllowSquashMerge:    repo.GetAllowSquashMerge(),                                                                                  // Whether squash merges are allowed
					AllowRebaseMerge:    repo.GetAllowRebaseMerge(),                                                                                  // Whether rebase merges are allowed
					AllowAutoMerge:      repo.GetAllowAutoMerge(),                                                                                    // Whether rebasing is allowed
					DeleteBranchOnMerge: repo.GetDeleteBranchOnMerge(),                                                                               // Topics associated with the repository
					Private:             repo.GetPrivate(),                                                                                           // Privacy status of the repository
					HasIssues:           repo.GetHasIssues(),                                                                                         // Whether the repository has issues
					HasWiki:             repo.GetHasWiki(),                                                                                           // Whether the repository has a wiki
					HasPages:            repo.GetHasPages(),                                                                                          // Whether the repository has pages
					IsTemplate:          repo.GetIsTemplate(),                                                                                        // Whether the repository is a template
					LicenseTemplate:     repo.GetLicenseTemplate(),                                                                                   // License template for the repository
					GitignoreTemplate:   repo.GetGitignoreTemplate(),                                                                                 // Gitignore template for the repository
					TeamId:              int64(repo.GetTeamID()),                                                                                     // Team ID associated with the repository
				}
				producer.SendRepository("repository", data) // Send repository details to Kafka topic
			}
		}
		return nil
	},
}

func init() {
	// Initialize Kafka producer
	var err error

	producer, err = kafka.NewProducer(brokerURL, schemaRegistryURL) // Initialize the Kafka producer with config
	if err != nil {
		panic(err)
	}

	fetchCmd.Flags().StringVarP(&profileType, "type", "t", "user", "Type of profile to fetch (user/org)")
	fetchCmd.Flags().StringVarP(&usernames, "usernames", "u", "yindia", "Comma-separated list of usernames")
	fetchCmd.Flags().StringVarP(&brokerURL, "broker-url", "b", "", "Kafka broker URL")
	fetchCmd.Flags().StringVarP(&schemaRegistryURL, "schema-registry-url", "s", "", "Schema registry URL")
	rootCmd.AddCommand(fetchCmd)
}
