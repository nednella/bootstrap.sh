---
name: react-best-practices
description: Performance and correctness rules for React 18 single-page apps built with Parcel or Vite, TypeScript and Tailwind. Use when writing, reviewing or refactoring React components in a client-side React app.
---

# React best practices

Rules for client-side React 18. They assume no server rendering, no React Server Components and no framework router. Read only the file for the category you need.

Measure before you optimise. Apply the memoisation rules only when a profile shows a problem.

## Categories, most impact first

| Category | File | Use when |
|---|---|---|
| Waterfalls | `rules/waterfalls.md` | Code awaits several async calls or loads data in nested components |
| Bundle size | `rules/bundle-size.md` | Adding a dependency, a heavy component or a third-party script |
| Data fetching | `rules/data-fetching.md` | Fetching in components, sharing listeners, using browser storage |
| Re-renders | `rules/rerender.md` | State, effects, props or hooks cause extra renders or bugs |
| Rendering | `rules/rendering.md` | Writing JSX, long lists, SVG, animations or script tags |
| JavaScript | `rules/javascript.md` | Hot loops, lookups, sorting, regular expressions, storage reads |
| Advanced | `rules/advanced.md` | Stable subscriptions with changing callbacks, one-time app setup |

## Not covered

Next.js, server components, server actions, edge or streaming rendering, and hydration. React 19 only APIs (`use`, `Activity`, `useEffectEvent`, async transitions) are also left out. Use refs where a rule would need them.

Adapted from vercel-labs/agent-skills react-best-practices (https://github.com/vercel-labs/agent-skills/tree/main/skills/react-best-practices). The upstream SKILL.md declares MIT but the repository has no LICENSE file, so the rules here are restated in new words, not copied.
