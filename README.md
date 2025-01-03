# Trout

Trout is a `Release to Production` tool for WordPress plugin updates. It operates in the same fashion as [Bowerbird](https://github.com/farghul/bowerbird.git), but with a focus on creating a WordPress production release via interaction with a Jira API. Trout is named after a popular catch and ***release*** fish.

![Trout](trout.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

An `env.json` file containing your API URL and Basic token to enable ticket creation:

``` json
{
    "token": "Email:Jira API token combination with Base 64 Encoding",
    "cloud": "Jira cloud issue base URL ex. https://jira.com/rest/api/latest/",
    "wordpress": "Path to the intended git repository containing composer-prod.json",
    "testing": "JQL Jira API search string to return a list of tickets with status 'In Progress'"
}
```

## Function

Trout searches the targeted Jira API for tickets marked as `In Progress` (aka Testing) for more than seven days. It then gathers the qualifying candidates and creates a new git branch named `update/[release]` where *release* is provided as an argument. Finally, it runs a series of `composer require` commands on the `composer-prod.json` file and prepares the new branch for a pull request.

## Build

From the root folder containing the `main.go` file, use the command that matches your environment:

### Windows & Mac:

``` zsh
go build -o [name] .
```

### Linux:

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

``` zsh
[program] [flag] [release name or number]
```

## Options

``` zsh
-r, --release   Run the main program
-h, --help      Help Information
-v, --version   Display Program Version
```

## Example

``` zsh
trout -r 88
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/trout/blob/main/LICENSE.md) and is part of the Public Domain.
