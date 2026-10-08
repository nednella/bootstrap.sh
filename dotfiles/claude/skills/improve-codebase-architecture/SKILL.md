---
name: improve-codebase-architecture
description: Scan a codebase for deepening opportunities, present them as a visual HTML report, then question the user through whichever one they pick. Use when the user asks to improve a codebase's architecture or find refactors that make it easier to test and navigate.
disable-model-invocation: true
---

# Improve codebase architecture

Find architectural friction and propose **deepening opportunities**: refactors that turn shallow modules into deep ones. The aim is code that is easier to test and easier for an AI to find its way around.

This skill uses a shared design vocabulary and the project's domain model:

- [CODEBASE-DESIGN.md](CODEBASE-DESIGN.md) defines the architecture terms (**module**, **interface**, **depth**, **seam**, **adapter**, **leverage**, **locality**) and the principles (the deletion test, "the interface is the test surface", "one adapter means a hypothetical seam, two means a real one"). Use these terms exactly in every suggestion. Do not drift into "component", "service", "API" or "boundary".
- `GLOSSARY.md`, if the project has one, names the domain concepts that make good seams. ADRs in `docs/adr/` record decisions this skill should not reopen.

The skill never changes code. It produces a report and a conversation.

## Process

### 1. Explore

**Decide where to look before you look.** Deepening a module pays off by making future changes easier, so weight the parts of the codebase that change often.

- If the user named a direction (a module, a subsystem, a pain point), take it and skip the next step.
- Otherwise, read a good stretch of `git log --oneline` to find the hot spots: the files and areas that keep coming up. Look there first. If the changes are scattered with no clear hot spot, widen the search.

Read `GLOSSARY.md` and any ADRs for the area first.

Then launch an agent with the Agent tool, `subagent_type: Explore`, to walk the code. Do not follow a fixed checklist. Explore freely and note where you hit friction:

- Where does understanding one concept mean jumping between many small modules?
- Where are modules **shallow**, with an interface nearly as complex as the implementation?
- Where were pure functions pulled out just for testing, while the real bugs hide in how they are called (no **locality**)?
- Where do tightly coupled modules leak across their seams?
- Which parts are untested, or hard to test through their current interface?

Apply the **deletion test** to anything that looks shallow: would deleting it concentrate complexity, or just move it? "Concentrates" is the signal you want.

### 2. Present candidates as an HTML report

Write one self-contained HTML file to the OS temp directory so nothing lands in the repo. Use `$TMPDIR`, falling back to `/tmp`, and name it `<tmpdir>/architecture-review-<timestamp>.html` so each run gets a fresh file. Open it for the user (`open <path>` on macOS, `xdg-open <path>` on Linux) and tell them the absolute path.

The report uses **Tailwind** and **Mermaid** from CDNs. Use Mermaid for graph-shaped relations (call graphs, dependencies, sequences), and hand-built divs or SVG for more editorial visuals. Give each candidate a **before and after diagram**.

Each candidate gets a card with:

- **Files**: the files and modules involved
- **Problem**: why the current design causes friction
- **Solution**: what would change, in plain English
- **Benefits**: in terms of locality and leverage, and how tests would improve
- **Before and after diagram**: side by side, showing the shallowness and the deepening
- **Strength**: `Strong`, `Worth exploring` or `Speculative`, shown as a badge

End with a **Top recommendation**: the candidate you would tackle first, and why.

Use `GLOSSARY.md` words for the domain and CODEBASE-DESIGN.md words for the architecture. If the glossary defines "Order", write "the Order intake module", not "the FooBarHandler" and not "the Order service".

**ADR conflicts**: if a candidate contradicts an ADR, only include it when the friction is real enough to reopen the ADR. Mark it on the card, e.g. "contradicts ADR-0007, but worth reopening because...". Do not list every refactor an ADR rules out.

See [HTML-REPORT.md](HTML-REPORT.md) for the scaffold, diagram patterns and styling.

Do not propose interfaces yet. Once the file is written, ask the user: "Which of these would you like to explore?"

### 3. Question the chosen candidate

When the user picks a candidate, walk its decisions with them: the constraints, the dependencies and their category (see [DEEPENING.md](DEEPENING.md)), the shape of the deepened module, what sits behind the seam, and which tests survive.

Work in rounds. Each round, ask every question whose prerequisites are already settled. Number them, give your recommended answer to each, and wait for the user's answers before the next round. Look up facts in the code yourself, and only put decisions to the user. Stop when no open decisions remain.

Update the domain model as decisions settle:

- **Naming a deepened module after a concept not in `GLOSSARY.md`?** Add the term. Create the file if it does not exist. Keep implementation details out of it.
- **Sharpening a fuzzy term in the conversation?** Update `GLOSSARY.md` there and then.
- **User rejects the candidate for a lasting reason?** Offer an ADR: "Want me to record this as an ADR so future reviews don't suggest it again?" Only offer when a future reviewer would need the reason to avoid suggesting it again. Skip passing reasons ("not worth it right now") and obvious ones.
- **Want to explore other interfaces for the deepened module?** Follow [DESIGN-IT-TWICE.md](DESIGN-IT-TWICE.md).

Adapted from mattpocock/skills (MIT).
