## qst view

Export the knowledge graph to a Mermaid.js diagram

### Synopsis

The view command reads the knowledge directory and exports its structure as a Mermaid.js graph. Prerequisite links are shown as solid arrows, and sub-guide relationships are shown as dotted arrows.

```
qst view [directory] [flags]
```

### Examples

```
# Output mermaid to stdout
qst view my_knowledge_graph_dir/

# Save to a markdown file
qst view . > graph.md
```

### Options

```
  -h, --help   help for view
```

### SEE ALSO

* [qst](qst.md)	 - A lossless compression and formatting tool for knowledge

