---
prerequisites:
    - Logical OR
    - Logical AND
sub_guides: []
scope: explanation
tags:
    - logic
    - discrete-math
---

## What is Exclusive OR (XOR)?

In propositional logic, **Exclusive OR** (often abbreviated as **XOR**) is a logical operation that is true if and only if exactly *one* of the propositions is true.

The Exclusive OR of two propositions $P$ and $Q$ is denoted as:

$$ P \oplus Q $$

(Read as *"P XOR Q"*).

Conceptually, XOR can be thought of as "P OR Q, but NOT both." 

### Truth Value Rule

A logical exclusive disjunction $P \oplus Q$ is **true** if and only if exactly **one** of the propositions is true. If both $P$ and $Q$ are true, or if both are false, the entire statement is false.

### Examples

- **Natural Language:**
  In everyday English, "or" is frequently used exclusively. 
  - Let $P$ = "You can have soup."
  - Let $Q$ = "You can have salad."
  - $P \oplus Q$ = "You can have soup or salad." *(Implies you cannot have both. If you take both, the waiter will charge you extra, making the original offer false!)*

- **Mathematical:**
  - $(3 < 5) \oplus (2 + 2 = 5)$ is **True** because exactly one component is true.
  - $(3 < 5) \oplus (2 + 2 = 4)$ is **False** because *both* components are true.
