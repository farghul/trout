package main

// Launch the program and execute according to the supplied flag
func main() {
	switch flag[1] {
	case "-h", "--help":
		help()
	case "-v", "--version":
		version()
	case "-r", "--release":
		serialize()
		trout = compiler()
		changedir()
		prepare()
		release = flag[2]
		packagist()
		journal("Branch " + branch + release + " for Production release " + release + " created.")
	default:
		alert("Unknown argument(s) -")
	}
}
