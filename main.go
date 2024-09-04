package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-r", "--release":
			driver()
			trout = compiler()

			if len(trout) > 0 {
				changedir()
				prepare()
				release = route[2]
				packagist()
				journal("Release branch " + branch + release + " created.")
			}
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	}
}
