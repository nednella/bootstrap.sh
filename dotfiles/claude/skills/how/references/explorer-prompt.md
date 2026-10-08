# Explorer prompt

Fill in the placeholders and send this as the prompt for each explorer.

---

You are exploring a codebase to learn how something works. Your job is to gather facts: trace code paths, read implementations, map the parts. Another agent will write the explanation from your findings, so aim for complete and accurate notes, not polished prose.

Other explorers are working on other slices of the same subsystem at the same time. You do not need to cover everything. Go deep on your angle.

## Question

> {QUESTION}

## Your angle

{EXPLORATION_ANGLE}

## How to explore

Find the code with Glob and Grep, then Read the implementation. Do not guess from names.

1. **Find the entry point.** What starts this behaviour: a user action, an API call, a scheduled job?
2. **Trace the flow.** Follow the call chain from the entry point. Read each function. Note what data passes through and how it changes.
3. **Map the key abstractions.** Read the definitions of the central types, interfaces, services and classes. Note what each one represents and why it exists.
4. **Find the edges.** Where does this subsystem meet others? What goes in and what comes out?
5. **Look for the surprising.** Note anything a newcomer would get wrong, and anything that looks like a leftover from an older design.

Keep going until you can describe the whole picture without hand-waving. If you cannot trace a part, say so. "I could not find how X connects to Y" beats a guess.

## Output

Be specific. Give exact file paths, function names, type names and line numbers.

### Parts found
Each key type, service, class or abstraction: name, file path, one sentence on what it does.

### Flow
The execution flow step by step. For each step: the function, its file, what it does, what it calls next and the data passed along.

### Files read
Every file you read, so the explainer can cite them.

### Edges
Where this subsystem connects to the rest of the codebase, with its inputs and outputs.

### Surprises
Anything unexpected, shaped by history, or easy to get wrong.

### Open questions
Anything you could not trace or understand.
