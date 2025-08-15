package main

const (
	bv     string = "1.0.0"
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	branch string = "release/"
	halt   string = "program halted "
	repos  string = "/data/automation/checkouts/"
	tokens string = "/data/automation/tokens/"
	config string = "desso-automation-conf/jsons/"
)

var (
	trout     []string
	plugin    string
	ticket    string
	release   string
	query     JQL
	jira      Jira
	token     Token
	bitbucket BitBucket
	jsons     = []string{repos + config + "bitbucket.json", repos + config + "jira.json", tokens + "tokens.json"}
)

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
	URL     string `json:"url"`
}

// Ticket holds the extracted data from the JQL queries
type JQL struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Status struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
			Updated     string `json:"updated"`
			Summary     string `json:"summary"`
			FixVersions []any  `json:"fixVersions"`
			Created     string `json:"created"`
		} `json:"fields"`
	} `json:"issues"`
}

// Tokens builds a list of jira tokens and api addresses
type Token struct {
	Bitbucket string `json:"bitbucket"`
	Jira      string `json:"jira"`
}
