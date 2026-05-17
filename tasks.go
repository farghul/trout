package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Test for an optional flag
func flag() string {
	var passed string
	if len(os.Args) < 3 {
		passed = "--short"
	} else {
		passed = os.Args[1]
	}
	return passed
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
		case 2:
			json.Unmarshal(data, &review)
		}
	}

	search, err := api(jira.Testing)
	inspect(err)
	json.Unmarshal(search, &query)
}

// Make the API call to Jira and return the response body
func api(criteria string) ([]byte, error) {
	baseURL := jira.URL + "search/jql?jql="

	fullURL := baseURL + criteria

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Basic "+jira.Token)
	req.Header.Set("Accept", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Iterate through the Args array and assign plugin and ticket values
func sift(box []string) {
	for i := 0; i < len(box); i++ {
		plugin = box[i]
		i++
		ticket = box[i]
		require()
		commit()
	}
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

// Alert the user if there are no tickets eligible for release
func proceed(task []string) {
	if len(task) == 0 {
		alert("No tickets eligible for release - ")
	}
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
