```go
func main() {
    d := Dev{
        name:       "nicolas goldack",
        age:        {{ .age }},
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
        stars:          {{.github.total_stars}},
        forks:          {{.github.total_forks}},
        followers:      {{.github.followers}},
    })
    d.GithubRepos(getRepos())
}

func getPopularRepos() {
    return []*Repo{
        &Repo{
            name: ngoldack/code-readme,
            stars: {{ index .github.repositories "code-readme" "stars" }},
            forks: {{ index .github.repositories "code-readme" "forks" }},
        },
        &Repo{
            name: ngoldack/Youtube-Channel-Badge,
            stars: {{ index .github.repositories "Youtube-Channel-Badge" "stars" }},
            forks: {{ index .github.repositories "Youtube-Channel-Badge" "forks" }},
        }
        &Repo{
            name: ngoldack/dicetrace,
            stars: {{ index .github.repositories "dicetrace" "stars" }},
            forks: {{ index .github.repositories "dicetrace" "forks" }},
        },
    }
}
```