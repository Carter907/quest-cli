---
prerequisites:
    - Propositions
sub_guides: []
clarity: detailed
scope: explanation
tags:
    - logic
    - discrete-math
---

## What is a Truth Table?

A **truth table** is a mathematical table used in logic to determine whether a compound proposition is true or false across all possible input values. It systematically lists all possible scenarios for the given variables.

### How to Construct a Truth Table

A standard truth table consists of:
- **Columns for variables**: The initial inputs, usually denoted by letters like $P$ and $Q$.
- **Rows for combinations**: Every possible combination of True (T) and False (F) for those inputs. For an expression with $n$ variables, there are $2^n$ rows.
- **Columns for outputs**: The resulting truth value of the logical expression evaluated step-by-step for each row.

### Visual Representation

Here is the truth table:

| $P$ | $Q$ | $\[another\ logical\ expression\]$ |
|:---:|:---:|:---:|
| T | T | **...** |
| T | F | **...** |
| F | T | **...** |
| F | F | **...** |

For each additional column added to the truth table, the previous values can be referenced to find the truth values of the rows for the next logical expression.
