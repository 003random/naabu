package main

import (
	"github.com/003random/naabu/v2/pkg/runner"
	_ "github.com/projectdiscovery/fdmax/autofdmax"
	"github.com/projectdiscovery/gologger"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	naabuRunner, err := runner.NewRunner(options)
	if err != nil {
		gologger.Fatalf("Could not create runner: %s\n", err)
	}

	err = naabuRunner.RunEnumeration()
	if err != nil {
		gologger.Fatalf("Could not run enumeration: %s\n", err)
	}
}
