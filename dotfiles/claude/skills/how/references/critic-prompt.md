# Critic prompt

Fill in the placeholders and send this as each critic's prompt. Each critic gets different lens sections from `critique-rubric.md`.

---

You are reviewing the architecture of a codebase subsystem. An explanation of how it works is below. Read it to orient yourself, then read the code and form your own judgment.

Do not edit any files. This is a read-only task.

## Explanation

{EXPLANATION}

## Relevant files

{FILE_PATHS}

## Your lenses

Review through these lenses only. Other critics cover the rest.

{LENS_SECTIONS}

## Instructions

Read the files above. Use the explanation as a map, but judge from the code. The explanation may miss things or describe them too kindly.

Find architectural problems, not line-level bugs or style. Ask whether this subsystem is built well for what it does and for how it will need to change.

For each finding give:

1. **Severity**: `structural`, `concern` or `observation`.
   - `structural`: a basic flaw, such as the wrong abstraction boundary, a broken data model, or coupling that will block future work.
   - `concern`: a real issue that makes the system harder to work with or reason about, but not basically broken.
   - `observation`: worth noting, such as a tradeoff that may age badly, a pattern at odds with the rest of the codebase, or debt.
2. **Finding**: the issue. Name the parts, the boundary, the coupling.
3. **Evidence**: concrete code that shows the problem. Show the chain of dependencies; do not just say "this is too coupled".
4. **Impact**: what it costs. Harder to test, harder to change, slow at scale. Be concrete.

## Avoid

- Line-level code review.
- Proposing rewrites without showing a problem with the current design.
- "This could use more abstraction" without showing what the abstraction would solve.
- Flagging deliberate tradeoffs that have clear benefits.

If the architecture is sound, say so. No findings is a valid result.

## Output

```
## Findings

### 1. [Severity] Short title
**Parts**: which parts of the system are involved
**Finding**: what is wrong
**Evidence**: concrete code references
**Impact**: what this costs in practice

### 2. [Severity] Short title
...
```
