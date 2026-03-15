package main

import (
	"flag"
	"fmt"
	"gostig/internal/cli"
)

func main() {
	acceptableActions := map[string]bool{
		"list":  true,
		"apply": true,
	}
	const (
		defaultAction = "list"
		actionUsage   = "The action that you want to make such as listing profiles, or applying a profile."
	)
	var action string
	if !acceptableActions[action] {
		flag.Usage()
		return
	}

	flag.StringVar(&action, "action", defaultAction, actionUsage)
	flag.StringVar(&action, "a", defaultAction, actionUsage)
	flag.Parse()

	switch action {
	case "list":
		cli.ListProfiles()
	case "apply": // Profile is determined programmatically.
		cli.ApplyProfile()

	}

	fmt.Print("listin those profiles")
	cli.ListProfiles()
}
