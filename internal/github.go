package internal

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func getClient(cfg *Config) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Github.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func getStats(cfg *Config) (map[string]interface{}, error) {
	client := getClient(cfg)
	stats := make(map[string]interface{})

	user, resp, err := client.Users.Get(context.TODO(), cfg.Github.Username)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	stats["followers"] = user.GetFollowers()
	stats["following"] = user.GetFollowing()
	stats["public_repos"] = user.GetPublicRepos()
	stats["public_gists"] = user.GetPublicGists()
	stats["created_at"] = user.GetCreatedAt()
	stats["updated_at"] = user.GetUpdatedAt()
	stats["bio"] = user.GetBio()
	stats["location"] = user.GetLocation()
	stats["email"] = user.GetEmail()
	stats["hireable"] = user.GetHireable()
	stats["blog"] = user.GetBlog()
	stats["twitter_username"] = user.GetTwitterUsername()
	stats["company"] = user.GetCompany()
	stats["name"] = user.GetName()
	stats["avatar_url"] = user.GetAvatarURL()
	stats["login"] = user.GetLogin()
	stats["type"] = user.GetType()
	stats["id"] = user.GetID()

	// Total Stars, Forks, Watchers, Issues
	repos, _, err := client.Repositories.List(context.TODO(), cfg.Github.Username, nil)
	if err != nil {
		return nil, err
	}
	var stars, forks, watchers, issuesOpen int
	for _, repo := range repos {
		stars += repo.GetStargazersCount()
		forks += repo.GetForksCount()
		watchers += repo.GetWatchersCount()
		issuesOpen += repo.GetOpenIssuesCount()
	}
	stats["total_stars"] = stars
	stats["total_forks"] = forks
	stats["total_watchers"] = watchers
	stats["total_issues_open"] = issuesOpen
	stats["total_repos"] = len(repos)
	stats["total_repos_private"] = user.GetTotalPrivateRepos()
	stats["total_repos_public"] = user.GetPublicRepos()

	// Repositories
	stats["repositories"] = make(map[string]interface{})
	for _, repo := range cfg.Github.Repositories {
		repo, _, err := client.Repositories.Get(context.TODO(), cfg.Github.Username, repo)
		if err != nil {
			return nil, err
		}

		stats["repositories"].(map[string]interface{})[repo.GetName()] = make(map[string]interface{})
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["forks"] = repo.GetForksCount()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["stars"] = repo.GetStargazersCount()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["watchers"] = repo.GetWatchersCount()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["issues"] = repo.GetOpenIssuesCount()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["size"] = repo.GetSize()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["language"] = repo.GetLanguage()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["license"] = repo.GetLicense().GetName()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["description"] = repo.GetDescription()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["created_at"] = repo.GetCreatedAt()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["updated_at"] = repo.GetUpdatedAt()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["pushed_at"] = repo.GetPushedAt()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["default_branch"] = repo.GetDefaultBranch()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["homepage"] = repo.GetHomepage()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["url"] = repo.GetURL()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["mirror_url"] = repo.GetMirrorURL()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["archived"] = repo.GetArchived()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["disabled"] = repo.GetDisabled()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["fork"] = repo.GetFork()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["has_issues"] = repo.GetHasIssues()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["has_projects"] = repo.GetHasProjects()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["has_wiki"] = repo.GetHasWiki()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["has_pages"] = repo.GetHasPages()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["has_downloads"] = repo.GetHasDownloads()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["private"] = repo.GetPrivate()
		stats["repositories"].(map[string]interface{})[repo.GetName()].(map[string]interface{})["is_template"] = repo.GetIsTemplate()

	}

	return stats, nil
}
