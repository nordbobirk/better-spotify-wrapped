package cli

import (
	"strings"
)

// Run is the entrypoint for the CLI.
func Run(args []string) {
	if len(args) <= 1 {
		Help()
		return
	}
	switch strings.ToLower(args[1]) {
	case "help":
		Help()
	case "parse":
		Parse(args[2:])
	case "favorite":
		Favorite(args[2:])
	case "top":
		Top(args[2:])
	case "hate":
		Hate(args[2:])
	default:
		Help()
	}
}
