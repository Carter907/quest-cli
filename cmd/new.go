package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new knowledge graph directory",
	Long:  "You start a new knowledge graph with the new command. You only have to specify the name of the knowledge graph; a directory will be created. A starter guide will be added in the directory as a template. If you don't specify a directory, the current directory will be used.",
	Example: `# Specify New Directory
qst new learn-cpp

# Use Current Directory
qst new`,
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) != 0 {
			dir = args[0]
		}

		err := os.MkdirAll(dir, 0o755)
		if err != nil {
			fmt.Printf("Failed to create knowledge graph: %v\n", err)
			os.Exit(1)
		}

		starterPath := filepath.Join(dir, "starter.md")
		if _, err := os.Stat(starterPath); os.IsNotExist(err) {
			starterContent := `---
prerequisites: []
sub_guides: []
clarity: strict
scope: definition
tags: ["example"]
---

## Starter Guide

This is an example guide. Replace this content with your own knowledge!
`
			err = os.WriteFile(starterPath, []byte(starterContent), 0o644)
			if err != nil {
				fmt.Printf("Failed to create starter guide: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Sample guide already exists, skipping write...\n")
		}

		metaFilePath := filepath.Join(dir, "manifest.yaml")
		if _, err := os.Stat(metaFilePath); os.IsNotExist(err) {
			metaFileContent := `title: Your New Knowledge
description: What can you say about this knowledge?
scopes:
  - definition
  - description
  - explanation
  - lesson
clarities:
  - strict
  - detailed
  - introductory
  - vague
relaxed_subguides: false
tours:
  - name: Tour 1
    guides:
      - Guide Filename 1
      - Guide Filename 2
  - name: Tour 2
    guides:
      -`

			err = os.WriteFile(metaFilePath, []byte(metaFileContent), 0o644)
			if err != nil {
				fmt.Printf("Failed to write metafile: %v\n", err)
				os.Exit(1)
			}

		} else {
			fmt.Printf("Metafile already exists, skipping write...\n")
		}
		fmt.Printf("Initialized knowledge graph in %s\n", dir)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
