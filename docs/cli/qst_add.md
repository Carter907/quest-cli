## qst add

Add a new guide to the knowledge graph.

### Synopsis

Add allows you to insert a guide into the knowledge graph by specifying its prerequisites, subguides, scope, and clarity. A new markdown file will be inserted into the directory

```
qst add [name] [flags]
```

### Examples

```
# Add a new definition
# missing value flags correspond to empty properties
qst add Exponent --scope definition --clarity strict
```

### Options

```
      --clarity string          Clarity of the guide (e.g. strict, vague)
  -h, --help                    help for add
  -i, --interactive             Interactive mode
      --prerequisites strings   List of prerequisites
      --scope string            Scope of the guide (e.g. definition, description)
      --subguides strings       List of subguides
      --tags strings            List of tags
```

### SEE ALSO

* [qst](qst.md)	 - A lossless compression and formatting tool for knowledge

