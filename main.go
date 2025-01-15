package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-r", "--release":
			serialize()
			trout = compiler()

			if len(trout) > 0 {
				changedir()
				prepare()
				release = route[2]
				packagist()
				journal("Branch " + branch + release + " for Production release " + release + " created.")
			}
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	}
}
