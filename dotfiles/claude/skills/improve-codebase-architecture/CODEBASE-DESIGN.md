# Codebase design

Design **deep modules**: a lot of behaviour behind a small interface, placed at a clean seam, testable through that interface. The aim is leverage for callers, locality for maintainers and testability for everyone.

## Glossary

Use these terms exactly. Do not swap in "component", "service", "API" or "boundary". Consistent language is the whole point.

**Module**: anything with an interface and an implementation, at any scale: a function, class, package or a slice across tiers. _Avoid_: unit, component, service.

**Interface**: everything a caller must know to use the module correctly. That covers the type signature, but also invariants, ordering rules, error modes, required configuration and performance. _Avoid_: API, signature (both cover only the type-level surface).

**Implementation**: the code inside a module. Distinct from **adapter**: a thing can be a small adapter with a large implementation (a Postgres repo) or a large adapter with a small implementation (an in-memory fake). Say "adapter" when the seam is the topic, "implementation" otherwise.

**Depth**: leverage at the interface. How much behaviour a caller or test can use per unit of interface it has to learn. A module is **deep** when a lot of behaviour sits behind a small interface, and **shallow** when the interface is nearly as complex as the implementation.

**Seam** (Michael Feathers): a place where you can change behaviour without editing that place. It is where a module's interface lives. Where to put the seam is its own decision, separate from what goes behind it. _Avoid_: boundary (overloaded by DDD's bounded context).

**Adapter**: a concrete thing that satisfies an interface at a seam. The word describes its role (the slot it fills), not what is inside.

**Leverage**: what callers get from depth. More capability per unit of interface learned. One implementation pays back across N call sites and M tests.

**Locality**: what maintainers get from depth. Changes, bugs, knowledge and checks gather in one place instead of spreading across callers. Fix once, fixed everywhere.

## Deep vs shallow

A **deep module** has a small interface and a large implementation:

```
┌─────────────────────┐
│   Small Interface   │  ← Few methods, simple params
├─────────────────────┤
│                     │
│  Deep Implementation│  ← Complex logic hidden
│                     │
└─────────────────────┘
```

A **shallow module** has a large interface and little implementation. Avoid it:

```
┌─────────────────────────────────┐
│       Large Interface           │  ← Many methods, complex params
├─────────────────────────────────┤
│  Thin Implementation            │  ← Just passes through
└─────────────────────────────────┘
```

When designing an interface, ask:

- Can I cut the number of methods?
- Can I simplify the parameters?
- Can I hide more complexity inside?

## Principles

- **Depth belongs to the interface, not the implementation.** A deep module can be built from small, mockable, swappable parts. They are just not part of the interface. A module can have **internal seams** (private, used by its own tests) as well as the **external seam** at its interface.
- **The deletion test.** Imagine deleting the module. If complexity vanishes, it was a pass-through. If complexity reappears across N callers, it was earning its keep.
- **The interface is the test surface.** Callers and tests cross the same seam. If you want to test past the interface, the module is probably the wrong shape.
- **One adapter means a hypothetical seam. Two adapters means a real one.** Do not add a seam unless something actually varies across it.

## Designing for testability

1. **Accept dependencies; do not create them.**

   ```typescript
   // Testable
   function processOrder(order, paymentGateway) {}

   // Hard to test
   function processOrder(order) {
     const gateway = new StripeGateway();
   }
   ```

2. **Return results; do not cause side effects.**

   ```typescript
   // Testable
   function calculateDiscount(cart): Discount {}

   // Hard to test
   function applyDiscount(cart): void {
     cart.total -= discount;
   }
   ```

3. **Keep the surface small.** Fewer methods need fewer tests. Fewer parameters need less test setup.

## Relationships

- A **module** has exactly one **interface**: the surface it shows callers and tests.
- **Depth** belongs to a **module** and is measured against its **interface**.
- A **seam** is where a **module**'s **interface** lives.
- An **adapter** sits at a **seam** and satisfies the **interface**.
- **Depth** gives **leverage** to callers and **locality** to maintainers.

## Rejected framings

- **Depth as the ratio of implementation lines to interface lines** (Ousterhout): this rewards padding the implementation. Use depth as leverage instead.
- **"Interface" as the TypeScript `interface` keyword or a class's public methods**: too narrow. Here the interface includes every fact a caller must know.
- **"Boundary"**: overloaded by DDD's bounded context. Say **seam** or **interface**.

## Going deeper

- [DEEPENING.md](DEEPENING.md): how to deepen a cluster given its dependencies, seam discipline, and replacing tests rather than layering them.
- [DESIGN-IT-TWICE.md](DESIGN-IT-TWICE.md): parallel agents design the interface several different ways, then you compare them on depth, locality and seam placement.
