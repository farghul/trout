package main

const (
	bv     string = "1.0.0"
	reset  string = "\033[0m"
	green  string = "\033[32m"
	yellow string = "\033[33m"
	red    string = "\033[41m"
	branch string = "release/"
	halt   string = "program halted "
	home   string = "/data/scripts/automation/programs/"
)

var (
	jira    Ticket
	plugin  string
	ticket  string
	release string
	trout   []string
	access  Atlassian
)

// Atlassian builds a list of jira tokens and api addresses
type Atlassian struct {
	BBA       string `json:"bba"`
	BitBucket string `json:"bitbucket"`
	Jira      string `json:"jira"`
	JQA       string `json:"jqa"`
	Testing   string `json:"testing"`
	Reviewer1 string `json:"reviewer1"`
	Reviewer2 string `json:"reviewer2"`
	WordPress string `json:"wordpress"`
}

// Ticket holds the extracted data from the JQL queries
type Ticket struct {
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
			Updated     string        `json:"updated"`
			Summary     string        `json:"summary"`
			FixVersions []interface{} `json:"fixVersions"`
			Created     string        `json:"created"`
		} `json:"fields"`
	} `json:"issues"`
}
