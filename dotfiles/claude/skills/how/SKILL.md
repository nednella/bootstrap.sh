---
name: how
description: Explain how a part of the codebase works by sending subagents to read the code, then writing up an architectural explanation. Use for "how does X work?" questions, and for critique when the user asks what is wrong with a subsystem's architecture.
---

# How

Answer "how does X work?" at the level a senior engineer needs to start working in a subsystem. Give them a working mental model, not annotated source code.

Two modes:

1. **Explain** (default): read the code and write an explanation.
2. **Critique**: explain first, then send critics to find architectural problems.

## Explain mode

### Step 1: Understand the question and judge its size

The user might ask about:

- a subsystem: "How does message virtualization work?"
- a feature flow: "How do we bill for on-demand usage?"
- a structure: "How is the auth service laid out?"
- a runtime trace: "What happens when a user sends a message?"

Work out the scope. If it is unclear, ask one short question before you explore.

Then pick a path:

- **Simple** (one module, a small utility, "how does function X work"): skip explorers. One explainer reads and writes in a single pass. Go to Step 2b.
- **Complex** (spans many files or services, crosses features, a full overview): send explorers in parallel, then hand off to the explainer. Go to Step 2a.

When in doubt, take the simple path. You can send explorers later if the explainer gets stuck.

### Step 2a: Explore (complex questions only)

Split the question into 2 to 4 angles, each a distinct slice so explorers do not repeat each other. For "how does message virtualization work?" you might split it into:

- the data model and state
- the rendering pipeline and DOM work
- the scroll and measurement code

Use 2 explorers for narrow questions and up to 4 for broad ones.

Launch all explorers in one message with the Agent tool, `subagent_type: Explore`. Each gets the prompt in `references/explorer-prompt.md` plus its own angle. Each explorer returns the parts it found, the flow it traced, the files it read and anything surprising. Overlap is fine. The explainer reconciles it.

Then go to Step 3.

### Step 2b: Direct explain (simple questions only)

Launch one agent with the Agent tool, `subagent_type: general-purpose`. It explores and writes the explanation in one pass. Build its prompt from `references/explainer-prompt.md`, leaving out the explorer findings. Tell it not to edit files.

Go to Step 4.

### Step 3: Synthesize (complex questions only)

When all explorers return, launch one agent with `subagent_type: general-purpose`. Build its prompt from `references/explainer-prompt.md` with every explorer's findings filled in. It merges overlaps, settles contradictions by checking the code, and writes one explanation.

### Step 4: Present

Show the user the explainer's output. You may edit lightly for clarity or add context from the conversation. Do not rewrite it.

For a large subsystem, also offer an HTML diagram of the flow.

### Output format

Adapt this to the question. Not every section fits every answer.

- **Overview**: one or two paragraphs on what it is, what it does and why it exists. A reader should be able to stop here and know whether to read on.
- **Key concepts**: the types, services or abstractions needed to follow the rest, each briefly defined.
- **How it works**: the core. What triggers it, what happens step by step, where data goes, where decisions happen. Prose, not pseudocode. Name files and functions so the reader can look.
- **Where things live**: the files and folders someone needs to start working here.
- **Gotchas**: surprises, sharp edges and the history behind code that looks odd.

## Critique mode

Use this when the user asks for architectural problems or improvements, not just understanding.

### Step 1: Explain first

Run the whole explain flow above. You cannot judge the design until you understand it.

### Step 2: Send critics

Launch up to three critics in one message, each with the Agent tool and `subagent_type: general-purpose`. Give each a different pair of lenses from `references/critique-rubric.md` so they do not reach the same findings:

| Critic | Lenses |
| ------ | ------ |
| A | Abstraction fit, complexity vs. value |
| B | Data model, boundary discipline |
| C | Evolution readiness, consistency |

Drop a critic whose lenses do not fit the subsystem. Build each prompt from `references/critic-prompt.md`. Each critic gets:

1. the explanation from Step 1, so it does not explore from scratch
2. the relevant file paths, so it can read the code
3. its two lens sections from the rubric

### Step 3: Lead judgment

Judge the findings as a practical lead. Do not just collect them.

- **Act on**: problems worth fixing now.
- **Consider**: real concerns where the cost and benefit are unclear.
- **Noted**: valid points of low priority.
- **Dismissed**: wrong, missing context, or a matter of taste.

Present the explanation first, then the verdict below it. The explanation must stand alone, so a reader who only wants to understand the system can skip the critique.

Adapted from poteto/how (MIT).
