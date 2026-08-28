---
prerequisites:
    - Truth Tables
sub_guides: []
clarity: detailed
scope: explanation
tags:
    - logic
    - discrete-math
---

## What is Logical Equivalence?

Two logical expressions are **logically equivalent** if they always have the exact same truth value, regardless of the truth values of their individual variables.

Logical equivalence is typically denoted by the double-arrow symbol ($\iff$) or the triple-bar symbol ($\equiv$). 

$$ A \iff B $$

This means "$A$ is logically equivalent to $B$" or "$A$ if and only if $B$."

### Using Truth Tables

Truth tables can be used to prove that two expressions are logically equivalent.

| $P$ | $Q$ | $P$ |
|:---:|:---:|:---:|
| T | T | **T** |
| T | F | **T** |
| F | T | **F** |
| F | F | **F** |

If two columns in a truth table have the same values row-for-row, then those expressions are equivalent. Here we added
P in another column to show this. Since $P \equiv P$, their columns should match.

When we introduce compound expressions using logical operators, you'll see that this idea becomes more significant as expressions get more complex.
