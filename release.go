package main

// A sequential list of tasks run to complete the program
func packagist() {
	checkout()
	execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "update", "--no-install")
	sift(trout)
	push()
	pullrequest()
}

// Switch to the desired branch, and pull any changes
func prepare() {
	execute("-e", "git", "checkout", "master")
	execute("-e", "git", "pull")
}

// Create a release branch if necessary
func checkout() {
	if exists(branch, release) {
		execute("-e", "git", "checkout", branch+release)
	} else {
		execute("-e", "git", "checkout", "-b", branch+release)
	}
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

// Run the appropriate composer require command based on the flag value
func require() {
	if edge() {
		execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "-W", "--no-install")
	} else {
		execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "--no-install")
	}
}

// Add and commit the updated code
func commit() {
	execute("-e", "git", "add", ".")
	execute("-e", "git", "commit", "-m", ticket+" "+plugin)
}

// Push the modified content to the git repository
func push() {
	execute("-e", "git", "push", "--set-upstream", "origin", branch+release)
}

// Create a pull request in BitBucket for the Production deployment release
func pullrequest() {
	execute("-e", "curl", "-L", "-X", "POST", "--url", bitbucket.URL+branch+release+"/pull-requests/", "--header", "Authorization: Basic "+bitbucket.Token, "--header", "Content-Type: application/json", "--data", "{'title': 'Release/"+release+"','source': {'branch': {'name': '"+branch+release+"'}}, 'destination': {'branch': {'name': 'master'}}, 'reviewers': [{'uuid': '"+bitbucket.Reviewers.One+"'}], 'close_source_branch': false}")
}
