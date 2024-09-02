package main

// A sequential list of tasks run to complete the program
func packagist() {
	checkout()
	execute("-e", "composer", "update", "--no-install")
	sift(trout)
	push()
}

// Run the appropriate composer require command based on the flag value
func require() {
	if edge() {
		execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "-W", "--no-install")
	} else {
		execute("-e", "env", "COMPOSER=composer-prod.json", "composer", "require", plugin, "--no-install")
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

// Create a release branch if necessary
func checkout() {
	if exists(relbranch, release) {
		execute("-e", "git", "switch", relbranch+release)
	} else {
		execute("-e", "git", "checkout", "-b", relbranch+release)
	}
}

// Add and commit the update
func commit() {
	execute("-e", "git", "add", ".")
	execute("-e", "git", "commit", "-m", ticket+" "+plugin)
}

// Push modified content to the git repository
func push() {
	execute("-e", "git", "push", "--set-upstream", "origin", relbranch+release)
}
