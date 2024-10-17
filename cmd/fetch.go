package cmd

import (
	"fmt"
	"strings"

	v1 "kafka/internal/gen/cloud/v1"
	"kafka/internal/github"
	"kafka/internal/kafka" // Import Kafka package

	"github.com/spf13/cobra"
)

var (
	profileType string          // "user" or "org"
	usernames   string          // Comma-separated usernames
	producer    *kafka.Producer // Kafka producer
)

var fetchCmd = &cobra.Command{
	Use:     "fetch",
	Short:   "Run the workflow orchestration server",
	Long:    `Start the workflow orchestration server and continuously stream task updates.`,
	Example: `  task fetch --type user --usernames user1,user2`,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := github.NewClient()

		for _, username := range strings.Split(usernames, ",") {
			if profileType == "user" {
				profile, err := client.UserProfile(strings.TrimSpace(username))
				if err != nil {
					return err
				}
				data := &v1.Profile{
					Login:     profile.Login,
					Id:        int32(profile.ID),
					AvatarUrl: profile.AvatarURL,
					HtmlUrl:   profile.HTMLURL,
					Type:      "user",
				}
				fmt.Printf("User Profile: %+v\n", profile)
				producer.Send("profiles", data) // Send user profile to Kafka topic
			} else if profileType == "org" {
				orgProfile, err := client.OrgProfile(strings.TrimSpace(username))
				if err != nil {
					return err
				}
				data := &v1.Profile{
					Login:     orgProfile.Login,
					Id:        int32(orgProfile.ID),
					AvatarUrl: orgProfile.AvatarURL,
					HtmlUrl:   orgProfile.HTMLURL,
					Type:      "org",
				}
				fmt.Printf("Organization Profile: %+v\n", orgProfile)
				producer.Send("profiles", data) // Send org profile to Kafka topic
			}

			// Fetch repositories
			repos, err := client.ListRepositories(strings.TrimSpace(username), 1)
			if err != nil {
				return err
			}
			fmt.Printf("Repositories for %s: %+v\n", username, repos)

			// Fetch details for each repository
			for _, repo := range repos {
				repoDetails, err := client.GetRepository(strings.TrimSpace(username), repo.Name)
				if err != nil {
					return err
				}
				fmt.Printf("Repository Details: %+v\n", repoDetails)
				data := &v1.Repository{
					Id:       int32(repoDetails.ID),
					Name:     repoDetails.Name,
					FullName: repoDetails.FullName,
					Private:  repoDetails.Private,
					Owner:    int32(repoDetails.ID),
				}
				producer.Send("repository", data) // Send repository details to Kafka topic

				// Fetch commits page by page
				page := 1
				for {
					commits, err := client.GetCommits(strings.TrimSpace(username), repo.Name, page)
					if err != nil {
						return err
					}
					if len(commits) == 0 {
						break
					}
					for _, commit := range commits {
						data := &v1.Commit{
							Sha: commit.SHA,
							Commit: &v1.Commit_CommitDetails{
								Author: &v1.Commit_CommitDetails_Author{
									Name:  commit.Commit.Author.Name,
									Email: commit.Commit.Author.Email,
									Date:  commit.Commit.Author.Date,
								},
								Message: commit.Commit.Message,
							},
						}
						producer.Send("commits", data) // Send commit details to Kafka topic
					}
					fmt.Printf("Commits for %s/%s (Page %d): %+v\n", username, repo.Name, page, commits)
					page++
				}
			}
		}
		return nil
	},
}

func init() {
	// Initialize Kafka producer
	var err error

	// Prompt for Kafka configuration
	var kafkaConfig string
	fmt.Print("Enter Kafka configuration (leave blank for default): ")
	fmt.Scanln(&kafkaConfig)

	if kafkaConfig == "" {
		kafkaConfig = "default_kafka_config" // Set your default Kafka config here
	}

	producer, err = kafka.NewProducer(kafkaConfig) // Initialize the Kafka producer with config
	if err != nil {
		panic(err)
	}

	fetchCmd.Flags().StringVarP(&profileType, "type", "t", "user", "Type of profile to fetch (user/org)")
	fetchCmd.Flags().StringVarP(&usernames, "usernames", "u", "", "Comma-separated list of usernames")
	rootCmd.AddCommand(fetchCmd)
}
