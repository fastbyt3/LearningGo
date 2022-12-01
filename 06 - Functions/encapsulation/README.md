# Encapsulation

- benefits:
    1. Clients cannot modify objects variables
    2. Hiding implementation details prevents clients from depending on things that might change
    3. Prevents clients from setting object's variables arbitarily

- `geometry.Path` has a sequence of points and no additional fields need to be added for that
- So we can safely def that as `Path []Points`
- But in case of `IntSet`, additional fields can be added or exisitng field can be modified depending upon use case
- This gives Encapsulation of `IntSet` more concrete reason


---
