# Trout

Trout is a `Release to Production` tool for WordPress plugin updates. It operates in the same fashion as [Bowerbird](https://github.com/farghul/bowerbird.git), but with a focus on creating a WordPress production release via interaction with a Jira API. Trout is named after a popular catch and ***release*** fish.

![Trout](trout.webp)

## Prerequisites

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

A `bitbucket.json` and `jira.json` file containing your API URLs and Basic tokens to enable authorized querying, ticket modification, and the creation of pull requests (see `jsons` folder for reference).

## Function

Trout searches the targeted Jira API for tickets marked as `In Progress` (aka Testing) for more than seven days. It then gathers the qualifying candidates and creates a new git branch named `update/[release]` where *release* is provided as an argument. Finally, it runs a series of `composer require` commands on the `composer-prod.json` file and creates a pull request for review.

## Build

Before building the application, change the value of the `resources` constant to reflect your environment:

``` go
resources string = "/data/automation/resources/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

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
-h, --help      Help Information
-r, --run       Run the main program
-v, --version   Display Program Version
```

## Example

``` zsh
trout -r 88
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/trout/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
