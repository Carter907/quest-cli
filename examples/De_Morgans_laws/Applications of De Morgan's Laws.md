---
prerequisites:
    - De Morgan's Laws
sub_guides: []
scope: explanation
tags:
    - logic
    - discrete-math
---

## Simplifying Symbolic Expressions

De Morgan's Laws are frequently used to simplify complex logical expressions. Here are a few examples:

- **Example 1**: Simplifying $\neg(A \land \neg B)$
  Using the first law, we distribute the negation and flip the operator:
  $\neg(A \land \neg B) \iff \neg A \lor \neg(\neg B) \iff \neg A \lor B$

- **Example 2**: Simplifying $\neg(\neg P \lor Q)$
  Using the second law, we distribute the negation and flip the operator:
  $\neg(\neg P \lor Q) \iff \neg(\neg P) \land \neg Q \iff P \land \neg Q$
