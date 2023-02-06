```go
func main() {
    d := Dev{
        name:       "nicolas goldack",
        age:        23,
        county:     c.Germany,
        hobbies:    []string{"coding", "gym", "games"},
        website:    URL{"http://ngoldack.de"},
    }

    d.CurrentEmployer(&Employer{
        name: "ARAG IT GmbH",
        since: 2019
    })
    d.CurrentPosition("full stack dev")

    d.SetSkills([]string{"go", "spring boot", "svelte"})

    d.GithubStats(&GithubStats{
        stars:          17,
        forks:          26,
        followers:      3,
    })
    d.GithubRepos(getRepos())
}

func getPopularRepos() {
    return []*Repo{
        &Repo{
            name: ngoldack/code-readme,
            stars:      0,
            forks:      0,
        },
        &Repo{
            name: ngoldack/Youtube-Channel-Badge,
            stars: 16,
            forks: 23,
        }
        &Repo{
            name: ngoldack/dicetrace,
            stars: 0,
            forks: 0,
        },
    }
}
```