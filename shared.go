package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var flag []string

// Test for an operational flag
func flags() []string {

	if len(os.Args) < 3 {
		flag[1] = "--zero"
	} else {
		flag[1] = os.Args[1]
	}
	return flag
}

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			json.Unmarshal(data, &bitbucket)
		case 1:
			json.Unmarshal(data, &jira)
		}
	}

	search := api(jira.Testing)
	json.Unmarshal(search, &query)
}

// Grab ticket information from the Jira API
func api(criteria string) []byte {
	result := execute("-c", "curl", "--request", "GET", "--url", jira.URL+"search?jql="+criteria, "--header", "Authorization: Basic "+jira.Token, "--header", "Accept: application/json")
	return result
}

// Build the list of candidates for production release
func compiler() []string {
	var candidate []string
	h, _ := time.ParseDuration("168h")
	for i := 0; i < len(query.Issues); i++ {
		if watchman(query.Issues[i].Fields.Updated) > h {
			candidate = append(candidate, query.Issues[i].Fields.Summary)
			candidate = append(candidate, query.Issues[i].Key)
		}
	}
	return candidate
}

// Determine how long a ticket status has been "In Progress"
func watchman(value string) time.Duration {
	stamp := strings.Replace(value, "-0800", "999999Z", 1)
	date, error := time.Parse(time.RFC3339Nano, stamp)
	inspect(error)
	waiting := time.Since(date)
	return waiting
}

// Confirm the current working directory is correct
func changedir() {
	os.Chdir(bitbucket.WordPress)
	var filePath string = "composer-prod.json"

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		alert("This is not the correct folder,")
	}
}

// Run a terminal command using flags to customize the output
func execute(variation, task string, args ...string) []byte {
	osCmd := exec.Command(task, args...)
	switch variation {
	case "-e":
		exec.Command(task, args...).CombinedOutput()
	case "-c":
		result, err := osCmd.Output()
		inspect(err)
		return result
	case "-v":
		osCmd.Stdout = os.Stdout
		osCmd.Stderr = os.Stderr
		err := osCmd.Run()
		inspect(err)
	}
	return nil
}

// Check to see if the latest release branch exists locally
func exists(prefix, tag string) bool {
	found := false
	b, _ := exec.Command("git", "branch").Output()
	if strings.Contains(string(b), prefix+tag) {
		found = true
	}
	return found
}

// Check for edge cases which require the -W flag
func edge() bool {
	found := false
	if strings.Contains(plugin, "roots/wordpress") {
		found = true
	}
	return found
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
