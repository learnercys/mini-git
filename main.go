package main

import (
	"fmt"
	"github.com/learnercys/mini-git/commands"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mini-git <command> [<args>]")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}

	switch command {
	case "init":
		if err := commands.Init(cwd); err != nil {
			fmt.Println("Error initializing repository:", err)
			os.Exit(1)
		}
	case "add":
		if len(args) < 1 {
			fmt.Println("Usage: mini-git add <file>")
			os.Exit(1)
		}
		if err := commands.Add(cwd, args[0]); err != nil {
			fmt.Println("Error adding file:", err)
			os.Exit(1)
		}

	case "commit":
		if len(args) < 1 {
			fmt.Println("Usage: mini-git commit <message>")
			os.Exit(1)
		}
		author := "John Doe <john@example.com> " // TODO: get author from config
		if err := commands.Commit(cwd, author, args[0]); err != nil {
			fmt.Println("Error committing changes:", err)
			os.Exit(1)
		}
	case "log":
		if err := commands.Log(cwd); err != nil {
			fmt.Println("Error showing commit log:", err)
			os.Exit(1)
		}
	case "status":
		if err := commands.Status(cwd); err != nil {
			fmt.Println("Error showing status:", err)
			os.Exit(1)
		}
	case "branch":
		if err := commands.Branch(cwd, args); err != nil {
			fmt.Println("Error managing branches:", err)
			os.Exit(1)
		}
	case "checkout":
		if err := commands.Checkout(cwd, args); err != nil {
			fmt.Println("Error checking out branch:", err)
			os.Exit(1)
		}
	case "merge":
		if err := commands.Merge(cwd, args); err != nil {
			fmt.Println("Error merging branches:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}
