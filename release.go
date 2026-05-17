package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// A sequential list of tasks run to complete the program
func packagist() {
	checkout()
	execute("env", []string{"COMPOSER=composer-prod.json", "composer", "update", "--no-install"}, ExecOptions{Stream: true})
	sift(trout)
	push()
	pullrequest()
}

// Switch to the development branch, and pull any changes
func prepare() {
	execute("git", []string{"checkout", "development"}, ExecOptions{Stream: true})
	execute("git", []string{"pull"}, ExecOptions{Stream: true})
}

// Create a release branch if necessary
func checkout() {
	if exists(branch, release) {
		execute("git", []string{"checkout", branch + release}, ExecOptions{Stream: true})
	} else {
		execute("git", []string{"checkout", "-b", branch + release}, ExecOptions{Stream: true})
	}
}

// Run the appropriate composer require command
func require() {
	if edge() {
		execute("composer", []string{"require", plugin, "-W", "--no-install"}, ExecOptions{Stream: true})
	} else {
		execute("composer", []string{"require", plugin, "--no-install"}, ExecOptions{Stream: true})
	}
}

// Stage all changes and commit the update
func commit() {
	execute("git", []string{"add", "."}, ExecOptions{Stream: true})
	execute("git", []string{"commit", "-m", ticket, "-m", "Install " + plugin}, ExecOptions{Stream: true})
}

// Push modified content to the git repository
func push() {
	execute("git", []string{"push"}, ExecOptions{Stream: true})
}

// Create a pull request in bitbucket
func pullrequest() {
	review.Title = "Release/" + release
	review.Description = "Production release for " + release
	review.Source.Branch.Name = branch + release
	review.Destination.Branch.Name = "master"

	jsonData, err := json.Marshal(review)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Create request
	req, err := http.NewRequest("POST", bitbucket.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Basic "+bitbucket.Token)
	req.Header.Set("Content-Type", "application/json")

	// HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Print result
	fmt.Printf("Response status: %s\n", resp.Status)
}
