# Design it twice

Use this when the user wants to explore other interfaces for a chosen deepening candidate. It follows Ousterhout's "Design It Twice": your first idea is unlikely to be the best.

Uses the terms in [CODEBASE-DESIGN.md](CODEBASE-DESIGN.md): **module**, **interface**, **seam**, **adapter**, **leverage**.

## Process

### 1. Frame the problem

Before launching agents, write the user an explanation of the problem for the chosen candidate:

- the constraints any new interface must meet
- the dependencies it would rely on, and their category (see [DEEPENING.md](DEEPENING.md))
- a rough code sketch that makes the constraints concrete. This is an illustration, not a proposal.

Show this to the user, then go straight to step 2. The user reads while the agents work.

### 2. Launch agents

Launch three or more agents in parallel with the Agent tool. Each must produce a very different interface for the deepened module.

Give each agent its own technical brief: file paths, coupling details, the dependency category from [DEEPENING.md](DEEPENING.md), and what sits behind the seam. This brief is separate from the explanation in step 1. Give each agent a different design constraint:

- Agent 1: "Minimize the interface: aim for one to three entry points. Maximize leverage per entry point."
- Agent 2: "Maximize flexibility: support many use cases and extension."
- Agent 3: "Optimize for the most common caller: make the default case trivial."
- Agent 4 (if it applies): "Design around ports and adapters for dependencies across the seam."

Include the terms from [CODEBASE-DESIGN.md](CODEBASE-DESIGN.md) and from `GLOSSARY.md` in the brief, so each agent names things in the architecture's and the domain's language.

Each agent returns:

1. The interface: types, methods and parameters, plus invariants, ordering and error modes
2. A usage example showing how callers use it
3. What the implementation hides behind the seam
4. The dependency strategy and adapters (see [DEEPENING.md](DEEPENING.md))
5. Trade-offs: where leverage is high and where it is thin

### 3. Present and compare

Present the designs one after another so the user can take each in, then compare them in prose. Contrast them by **depth** (leverage at the interface), **locality** (where change gathers) and **seam placement**.

Then give your own recommendation: which design is strongest and why. If parts of different designs would combine well, propose a hybrid. Take a clear position. The user wants a strong read, not a menu.
