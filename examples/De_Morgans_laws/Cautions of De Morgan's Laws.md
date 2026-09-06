---
prerequisites:
    - De Morgan's Laws
sub_guides: []
scope: explanation
tags:
    - logic
    - discrete-math
---

## 1. Forgetting to Flip the Operator

When first learning De Morgan's Laws, the most common mathematical mistake is distributing the negation to the variables but forgetting to flip the logical operator between them. 

For example, when simplifying the negation of a conjunction, $\neg(A \land B)$, it is incorrect to write $\neg A \land \neg B$. You must flip the "AND" ($\land$) to an "OR" ($\lor$):

**Incorrect:** $\neg(A \land B) \iff \neg A \land \neg B$
**Correct:** $\neg(A \land B) \iff \neg A \lor \neg B$

## 2. Grammatical Ambiguity in English

The correct negation of the sentence "Jim is tall and thin" is "Jim is not tall or Jim is not thin". 

However, we should be careful because in everyday language, people often just say:

> "Jim is not tall and thin."

This phrasing appears logically invalid if interpreted strictly as $\neg A \land \neg B$. But because of the grammatical rules of English, this sentence is generally understood by native speakers to be logically equivalent to the correct De Morgan's expansion. 

This happens because, in formal logic, operators like "and" only apply to complete propositions, whereas natural language allows us to casually group adjectives. Therefore, it is critical to carefully distinguish between formal mathematical logic and casual grammatical context.
