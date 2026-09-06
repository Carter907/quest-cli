## qst link

Link a subguide to an existing guide.

### Synopsis

Modifies an existing parent guide's frontmatter to add a new subguide relation with adherence and segment.

```
qst link [parent-guide] [flags]
```

### Examples

```
# Link a subguide interactively
qst link Exponent -i

# Link a subguide via flags
qst link Exponent --guide Power --adherence detailed --segment 1-5
```

### Options

```
      --adherence string          Adherence of the subguide relation
  -d, --dir string              Knowledge graph directory (default ".")
      --guide string            Name of the subguide to link
  -h, --help                    help for link
  -i, --interactive             Interactive mode
      --segment string          Segment of the subguide relation
```

### SEE ALSO

* [qst](qst.md)	 - A lossless compression and formatting tool for knowledge
