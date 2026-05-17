package main

// BitBucket builds a list of BitBucket tokens and api addresses
type BitBucket struct {
	URL       string `json:"url"`
	UUID      string `json:"uuid"`
	WordPress string `json:"wordpress"`
	Reviewers struct {
		One   string `json:"one"`
		Two   string `json:"two"`
		Three string `json:"three"`
	}
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

type Color string

const (
	Reset         = "\033[0m"
	Red    Color  = "\033[31m"
	Green  Color  = "\033[32m"
	Yellow Color  = "\033[33m"
	BGRed  Color  = "\033[41m"
	bv     string = "1.0.0"
	branch string = "release/"
	halt   string = "program halted "
	config string = "/data/automation/jsons/"
	repos  string = "/data/automation/repos/"
	tokens string = "/data/automation/tokens/"
)

var (
	query     JQL
	jira      Jira
	plugin    string
	ticket    string
	release   string
	trout     []string
	bitbucket BitBucket
	jsons     = []string{repos + config + "bitbucket.json", repos + config + "jira.json", tokens + "tokens.json"}
)
