# Rendering

## Do not render stray zeros

`{count && <Badge />}` renders `0` when `count` is `0`. Compare explicitly or use a ternary.

```tsx
{count > 0 ? <span className="badge">{count}</span> : null}
```

## Hoist static JSX

JSX that never changes can live at module level, so React does not rebuild it on each render. This helps most with large static SVG. The gain is small for most elements, so do it when a profile points there.

```tsx
const loadingSkeleton = <div className="h-20 animate-pulse bg-gray-200" />
```

## Use content-visibility on long lists

`content-visibility: auto` lets the browser skip layout and paint for off-screen items. Give each item an estimated size so the scrollbar stays stable.

```css
.message-item {
  content-visibility: auto;
  contain-intrinsic-size: 0 80px;
}
```

In Tailwind, use `[content-visibility:auto]` and `[contain-intrinsic-size:0_80px]`. For very long lists, use a virtual list instead.

## Animate a wrapper, not the SVG

Some browsers do not hardware-accelerate CSS animations on `<svg>` elements. Put the animation class on a wrapping `<div>`.

```tsx
<div className="animate-spin">
  <svg width="24" height="24" viewBox="0 0 24 24">
    <circle cx="12" cy="12" r="10" stroke="currentColor" />
  </svg>
</div>
```

This applies to `transform`, `opacity` and similar properties.

## Trim SVG precision

Extra decimals in path data add bytes and no visible detail. One decimal is usually enough for a small viewBox.

```bash
npx svgo --precision=1 --multipass icon.svg
```

## Load scripts without blocking

A plain `<script src>` in `index.html` blocks parsing. Use `async` for independent scripts such as analytics and `defer` for scripts that need the DOM or run in order. Module scripts that Vite and Parcel inject are already deferred.

```html
<script src="https://example.com/analytics.js" async></script>
<script src="/scripts/utils.js" defer></script>
```
