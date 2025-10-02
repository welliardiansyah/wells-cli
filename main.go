package main

import "github.com/welliardiansyah/wells-cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
