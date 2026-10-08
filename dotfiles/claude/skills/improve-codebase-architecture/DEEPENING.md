# Deepening

How to deepen a cluster of shallow modules safely, given its dependencies. Uses the terms in [CODEBASE-DESIGN.md](CODEBASE-DESIGN.md): **module**, **interface**, **seam**, **adapter**.

## Dependency categories

Sort a candidate's dependencies into these categories. The category decides how you test the deepened module across its seam.

### 1. In-process

Pure computation, in-memory state, no I/O. Always deepenable: merge the modules and test through the new interface. No adapter needed.

### 2. Local-substitutable

Dependencies with a local stand-in for tests, such as PGLite for Postgres or an in-memory filesystem. Deepenable if the stand-in exists. Tests run the deepened module against the stand-in. The seam is internal, with no port at the module's external interface.

### 3. Remote but owned (ports and adapters)

Your own services across a network, such as microservices or internal APIs. Define a **port** (an interface) at the seam. The deep module owns the logic, and the transport is injected as an **adapter**. Tests use an in-memory adapter. Production uses an HTTP, gRPC or queue adapter.

Recommendation shape: "Define a port at the seam, with an HTTP adapter for production and an in-memory adapter for tests, so the logic sits in one deep module even though it is deployed across a network."

### 4. True external (mock)

Third-party services you do not control, such as Stripe or Twilio. The deepened module takes the external dependency as an injected port, and tests supply a mock adapter.

## Seam discipline

- **One adapter means a hypothetical seam. Two adapters means a real one.** Do not add a port unless at least two adapters are justified, usually production and test. A seam with one adapter is just indirection.
- **Internal seams vs external seams.** A deep module can have internal seams, private to its implementation and used by its own tests, as well as the external seam at its interface. Do not expose internal seams through the interface just because tests use them.

## Testing: replace, do not layer

- Once tests exist at the deepened module's interface, the old unit tests on the shallow modules are waste. Delete them.
- Write new tests at the deepened module's interface. The **interface is the test surface**.
- Tests assert on outcomes seen through the interface, not on internal state.
- Tests should survive internal refactors because they describe behaviour, not implementation. If a test must change when the implementation changes, it is testing past the interface.
