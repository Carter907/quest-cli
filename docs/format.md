# File Format

Knowledge graphs are archived into `.kng` files which are zip archives containing markdown files (`.md`) denoted as "Guides". These guides have a particular format that makes them ideal for housing and structuring knowledge in a Directed Acyclic Graph (DAG).

## Guide Metadata

These guides require a YAML frontmatter that carries their metadata. When formatting these markdown files you should make sure all fields specified below are included.

> [!NOTE]
> When creating metadata about language itself, it's important to realize that we are rubbing up against philosophical and linguistic barriers. There is no universally accepted way of measuring "scope" or "clarity" in a piece of text. Therefore, defining these terms qualitatively gives us some breathing room for the fuzziness in their definitions.

**Properties**:

- `prerequisites` - A list of required guides. **Strict Rule:** Horizontal edges in the knowledge graph can *only* exist between guides of the exact identical `scope` (horizontal relationship).
- `sub_guides` - An optional list of encompassed guides. **Strict Rule:** These must be exactly one scope level smaller than the current guide, unless `relaxed_subguides: true` is set in the `manifest.yaml` which allows any smaller scope.
- `clarity` - This is a **functional text constraint**, not just a label. It dictates exactly how the `sub_guides` must be tangibly represented and rewritten inside the body text of the current guide.
  - **strict**: The author must explicitly use the exact or near-exact content of the sub-guides within the text.
  - **detailed**: The author must explicitly write closely summarized versions of the sub-guides within the text.
  - **introductory**: The author only needs to generalize the sub-guides, pointing to the concept without exact reproduction.
  - **vague**: The author may reference the sub-guides loosely, with the least exact reproduction of content.
- `scope` - How much content is covered in a guide; how many concepts or things were explained. Scope is qualitative:
  - **definition**: Smallest scope (singular term). *Example: "Exponent"*
  - **description**: Smaller scope, but slightly larger than a definition, focusing heavily on providing examples and comparisons. *Example: "Graphs of Exponential Function"*
  - **explanation**: Medium guide that explicitly relies on multiple descriptions to explain a mechanic. *Example: "Real-valued Functions"*
  - **lesson**: Largest content type that aggregates and teaches multiple explanations. *Example: "Understanding Graphing in Algebra"*

**Tags**:

Tags are subjects or topic descriptors that summarize the content of the entire guide.

**Title**:

The title of a guide is defined by its file name.
