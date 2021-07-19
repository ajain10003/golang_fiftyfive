package main

import (
	"fiftyfive/cmd"
	"fiftyfive/pkg/logex"
)

// Name of the application
const serviceName = "Supermarket"

// Version is set during build via --id flag parameter
var Version = "untagged Build"

func main() {
	// setup logger
	logger, flush := logex.SetupAndBuild(serviceName)
	defer flush()

	// call to execute
	cmd.Execute(logger, serviceName, Version)
}
