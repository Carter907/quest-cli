package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Carter907/quest-cli/internal/graph"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [directory]",
	Short: "Export the knowledge graph to a Mermaid.js diagram",
	Long:  `The view command reads the knowledge directory and exports its structure as a Mermaid.js graph. Prerequisite links are shown as solid arrows, and sub-guide relationships are shown as dotted arrows.`,
	Example: `# Output mermaid to stdout
qst view my_knowledge_graph_dir/

# Save to a markdown file
qst view . > graph.md`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		guides, err := graph.ParseGraph(dir)
		if err != nil {
			fmt.Printf("Error parsing graph: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(generateMermaid(guides))
	},
}

func generateMermaid(guides map[string]graph.Guide) string {
	var sb strings.Builder
	sb.WriteString("```mermaid\ngraph TD\n")

	// Sort guides for deterministic output
	var ids []string
	for id := range guides {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	// Declare all nodes first (ensures nodes with no connections are still visible)
	for _, id := range ids {
		cleanID := strings.ReplaceAll(id, "-", "_")
		fmt.Fprintf(&sb, "    %s[\"%s\"]\n", cleanID, id)
	}

	// Declare all edges
	for _, id := range ids {
		guide := guides[id]
		cleanID := strings.ReplaceAll(id, "-", "_")

		// Prerequisites (Solid links: Prereq -> Current)
		for _, prereq := range guide.Metadata.Prerequisites {
			cleanPrereq := strings.ReplaceAll(prereq, "-", "_")
			fmt.Fprintf(&sb, "    %s --> %s\n", cleanPrereq, cleanID)
		}

		// Subguides (Dotted links: Current -> Sub)
		for _, sub := range guide.Metadata.SubGuides {
			cleanSub := strings.ReplaceAll(sub.Guide, "-", "_")
			fmt.Fprintf(&sb, "    %s -.-> %s\n", cleanID, cleanSub)
		}
	}

	sb.WriteString("```")
	return sb.String()
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
