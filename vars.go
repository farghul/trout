package main

// Bitbucket builds a list of Bitbucket tokens and api addresses
type Bitbucket struct {
	URL       string `json:"url"`
	UUID      string `json:"uuid"`
	Token     string `json:"token"`
	WordPress string `json:"wordpress"`
}

// Jira builds a list of jira tokens and api addresses
type Jira struct {
	Testing string `json:"testing"`
	Basic   string `json:"basic"`
	Token   string `json:"token"`
	URL     string `json:"url"`
}

// JQL holds the extracted data from the JQL queries
type JQL struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Updated string `json:"updated"`
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

// Struct for PR payload
type PullRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	Source struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"source"`

	Destination struct {
		Branch struct {
			Name string `json:"name"`
		} `json:"branch"`
	} `json:"destination"`

	Reviewers []struct {
		UUID string `json:"uuid"`
	} `json:"reviewers"`

	CloseSourceBranch bool `json:"close_source_branch"`
}

type Color string

const (
	Reset         = "\033[0m"
	Red    Color  = "\033[31m"
	Green  Color  = "\033[32m"
	Yellow Color  = "\033[33m"
	BGRed  Color  = "\033[41m"
	bv     string = "1.1.0"
	branch string = "release/"
	halt   string = "program halted "
	config string = "/data/automation/jsons/"
)

var (
	query     JQL
	jira      Jira
	plugin    string
	ticket    string
	release   string
	trout     []string
	bitbucket Bitbucket
	review    PullRequest
	jsons     = []string{config + "bitbucket.json", config + "jira.json", config + "pullrequest.json"}
)
