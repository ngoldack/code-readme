# code-readme

Populate a code readme with actual values.

## Description

This project is a github action and a cli tool to populate a code readme with actual values. It uses the github api to get the values.

## Table of Contents

| Section | Description |
| --- | --- |
| [Examples](#examples) |  |
| [Features](#features) | List of all available features |
| [Using the github action](#using-the-github-action) | How to use the github action |
| [Using the cli](#using-the-cli) | How to use the cli |


## Examples

- This repo has {{ index .github.repositories "code-readme" "stars" }} stars.
- This repo has {{ index .github.repositories "code-readme" "forks" }} forks.
- This repo has {{ index .github.repositories "code-readme" "watchers" }} watchers.

See my personal [README.md](https://github.com/ngoldack/ngoldack/blob/main/README.md).
For more examples see the [examples directory](https://github.com/ngoldack/code-readme/tree/main/example).

## Features

| Feature | Description | Usage | Example |
| --- | --- | --- | --- |
| Age | Calculates the age of the developer | `{{ .age }}` | `{{ .age }}` |
| Github Stars | The total amount of stars of all repositories | `{{ .github.total_stars }}` | `{{ .github.total_stars }}` |
| Github Forks | The total amount of forks of all repositories | `{{ .github.total_forks }}` | `{{ .github.total_forks }}` |
| Github Followers | The total amount of followers | `{{ .github.followers }}` | `{{ .github.followers }}` |
| Github Repository Stars | Amount of stars of the given repository | `{{ index .github.repositories "[REPOSITORY_NAME]" "stars" }}` | `{{ index .github.repositories "code-readme" "stars" }}` |
| Github Repository Forks | Amount of forks of the given repository  | `{{ index .github.repositories "[REPOSITORY_NAME]" "forks" }}` | `{{ index .github.repositories "code-readme" "forks" }}` |
| Github Repository Watchers | Amount of stars of the given repository | `{{ index .github.repositories "[REPOSITORY_NAME]" "forks" }}` | `{{ index .github.repositories "code-readme" "forks" }}` |

Replace `[VARIABLE]` with a valid variable including the squared brackets.

## Usage

### Using the github action

You can use this project in a github action. The action will generate a readme for your repository.

### Using the cli

You can use the cli to generate a readme for your repository.


## Configuration

### Method A: Use a config file (recommended)

### 

---

## Thanks to the following projects

- [github-readme-stats](https://github.com/anuraghazra/github-readme-stats) - inspiration
- [cobra](https://github.com/spf13/cobra) - cli framework
- [viper](https://github.com/spf13/viper) - config management
- [pterm](https://github.com/pterm/pterm) - cli output