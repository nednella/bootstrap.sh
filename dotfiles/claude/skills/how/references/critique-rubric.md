# Critique rubric

Six lenses, paired so each critic gets two:

- Critic A: abstraction fit, complexity vs. value
- Critic B: data model, boundary discipline
- Critic C: evolution readiness, consistency

Paste a critic's two sections into its prompt. Not every lens fits every subsystem.

## Abstraction fit

Do the abstractions earn their place?

- Does each abstraction stand for a real concept, or is it a layer added "in case we need it"?
- Are the boundaries in the right place? Do they separate things that change on their own?
- Do two parts share implementation details they should not need to know?
- Is business logic tangled with framework wiring, or kept apart?

Too much abstraction is as bad as too little. A flat, simple design is fine for a simple domain.

## Complexity vs. value

Is the complexity spent well?

- Where is the complexity? In the parts that need it (core logic, tricky invariants) or in accidental places (boilerplate, needless indirection, configuration)?
- Is there a simpler way to get the same behaviour?
- Does every part earn its place, or are some left over from an earlier design?

## Data model

Do the data structures fit how the data is used?

- Are the models built for how data is read and written, or for how someone first pictured it?
- Does code keep reshaping data because the model does not match how it is used?
- Are the types honest? Do they describe the data as it is at runtime, or claim more structure than exists?

## Boundary discipline

Are the system's boundaries clean and well placed?

- Is validation done at entry points, or scattered through internal code?
- Are errors handled at boundaries and passed on cleanly, or caught and re-thrown at every layer?
- Does data cross boundaries in well-typed shapes, or as bags of optional fields?
- Can this subsystem be tested alone, or does it need the whole system running?

## Evolution readiness

How well will this design handle likely changes?

- If the most likely next requirement landed tomorrow, how much would change? One file, or everything?
- Which hardcoded assumptions would need to be relaxed?
- Does the design look bolted on, or as if it was always part of the plan?
- Are there legacy paths kept for compatibility that nothing depends on?

Do not penalize the design for missing hypothetical changes. Focus on changes the codebase's direction makes plausible.

## Consistency

Does this subsystem follow the patterns used elsewhere in the codebase?

- Are similar problems solved here the same way as elsewhere, or does this area invent its own patterns?
- If the patterns differ, is there a good reason, or did they drift apart?
- A difference is not bad by itself. An unexplained one is a maintenance cost.
