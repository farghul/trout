package main

// A sequential list of tasks run to complete the program
func packagist() {
	checkout()
	execute("env", []string{"COMPOSER=composer-prod.json", "composer", "update", "--no-install"}, ExecOptions{Stream: true})
	sift(trout)
	push()
	// pullrequest()
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

// Add and commit the update
func commit() {
	execute("git", []string{"add", "."}, ExecOptions{Stream: true})
	execute("git", []string{"commit", "-m", ticket, "-m", "Install " + plugin}, ExecOptions{Stream: true})
}

// Push modified content to the git repository
func push() {
	execute("git", []string{"push"}, ExecOptions{Stream: true})
}

/* Needs more work
// Create a pull request in BitBucket for the Production deployment release
func pullrequest() {
	execute("-v", "curl", "-L", "-X", "POST", "--url", bitbucket.URL+branch+release+"/pull-requests/", "--header", "Authorization: Basic "+token.Bitbucket, "--header", "Content-Type: application/json", "--data", "{'title': 'Release/"+release+"','source': {'branch': {'name': '"+branch+release+"'}}, 'destination': {'branch': {'name': 'master'}}, 'reviewers': [{'uuid': '"+bitbucket.Reviewers.One+"'}], 'close_source_branch': false}")
}
*/
