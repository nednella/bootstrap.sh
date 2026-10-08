# JavaScript

Small gains that matter in hot paths: render bodies, event handlers and loops over large lists. Skip them for arrays of ten items.

## Arrays and lookups

- Build a `Map` once when you call `.find()` by the same key many times.

  ```ts
  const userById = new Map(users.map((u) => [u.id, u]))
  return orders.map((o) => ({ ...o, user: userById.get(o.userId) }))
  ```

- Use a `Set` for repeated membership checks: `allowed.has(id)` instead of `list.includes(id)`.
- Combine several `.filter()` calls over one array into a single loop that fills each result.
- Replace `.map(...).filter(Boolean)` with `.flatMap((x) => (ok(x) ? [f(x)] : []))`.
- Find a min or max with one loop, not a sort.
- Check lengths first when comparing arrays. Different lengths mean different arrays, so skip the sort or deep compare.
- Use `.toSorted()`, `.toReversed()` and `.toSpliced()` instead of mutating methods. `.sort()` on a prop or state array mutates it and breaks React. Where `toSorted` is unavailable, use `[...items].sort(...)`.
- Return early from functions and loops once the answer is known.
- Read a deep property once before a loop, not inside it.

## Cache repeated work

- Store results of a pure function called many times with the same input in a module-level `Map`.
- Reads from `localStorage`, `sessionStorage` and `document.cookie` are synchronous and slow. Cache them in a `Map`, update the cache on writes, and clear an entry on the `storage` event so other tabs stay in sync.
- Create a `RegExp` once at module level, or in `useMemo` when it depends on input. A regex with the `g` flag keeps `lastIndex` between calls, so do not share one across `.test()` calls without resetting it.

## Browser work

- Do not interleave style writes with layout reads (`offsetWidth`, `getBoundingClientRect()`, `getComputedStyle()`). Each read after a write forces a reflow. Batch the writes, then read once. Prefer toggling a CSS class over many inline styles.
- Run non-critical work such as analytics and prefetching in `requestIdleCallback`. Pass `{ timeout }` when it must run eventually, and check browser support or add a `setTimeout` fallback.

  ```ts
  requestIdleCallback(() => analytics.track('search', { query }), { timeout: 2000 })
  ```
