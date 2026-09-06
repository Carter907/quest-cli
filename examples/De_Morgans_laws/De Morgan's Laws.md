---
prerequisites:
    - Logical Equivalence
    - Logical AND
    - Logical OR
    - Negation
sub_guides: []
scope: explanation
tags:
    - logic
    - discrete-math
---

## Symbolic Representation

De Morgan's Laws can be expressed in propositional logic as:

1. **Negation of Conjunction:**
   $$ \neg (A \land B) \iff (\neg A \lor \neg B) $$
   "Not (A and B) is the same as (not A) or (not B)."

2. **Negation of Disjunction:**
   $$ \neg (A \lor B) \iff (\neg A \land \neg B) $$
   "Not (A or B) is the same as (not A) and (not B)."

## Truth Table

You can prove that these logical expressions are equivalent using truth tables:

### First Law: $\neg(A \land B) \iff \neg A \lor \neg B$

| $A$ | $B$ | $A \land B$ | $\neg(A \land B)$ | $\neg A$ | $\neg B$ | $\neg A \lor \neg B$ |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| T | T | T | **F** | F | F | **F** |
| T | F | F | **T** | F | T | **T** |
| F | T | F | **T** | T | F | **T** |
| F | F | F | **T** | T | T | **T** |


You know they have the same truth value because the columns of $\neg(A \land B)$ and $\neg A \lor \neg B$ are equal.
The same is true for the second expression:

### Second Law: $\neg(A \lor B) \iff \neg A \land \neg B$

| $A$ | $B$ | $A \lor B$ | $\neg(A \lor B)$ | $\neg A$ | $\neg B$ | $\neg A \land \neg B$ |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| T | T | T | **F** | F | F | **F** |
| T | F | T | **F** | F | T | **F** |
| F | T | T | **F** | T | F | **F** |
| F | F | F | **T** | T | T | **T** |
