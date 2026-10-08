# Re-renders and effects

Several rules here involve `memo`, `useMemo` and `useCallback`. Add those only after a profile shows a cost. Rules that remove state or effects are always safe.

## Derive values during render

If a value follows from props or state, compute it in the render body. Do not copy it into state and sync it in an effect.

```tsx
const [firstName, setFirstName] = useState('First')
const [lastName, setLastName] = useState('Last')
const fullName = `${firstName} ${lastName}`
```

To reset state when a prop changes, give the component a `key` instead of setting state in an effect.

## Run user actions in event handlers

Do not model a click or submit as `setSubmitted(true)` plus an effect. The effect re-runs on unrelated changes and can fire twice. Put the work in the handler.

```tsx
function Form() {
  const theme = useContext(ThemeContext)

  function handleSubmit() {
    post('/api/register')
    showToast('Registered', theme)
  }

  return <button onClick={handleSubmit}>Submit</button>
}
```

## Never define components inside components

A component defined in a render body is a new type on every render. React remounts it, which loses state, focus and scroll position. Define it at module level and pass props.

Symptoms: an input loses focus on each keystroke, an animation restarts, an effect re-runs on every parent render.

## Narrow effect dependencies

Depend on primitives, not whole objects.

```tsx
useEffect(() => {
  log(user.id)
}, [user.id])
```

For thresholds, compute the boolean first and depend on it.

```tsx
const isMobile = width < 768
useEffect(() => {
  if (isMobile) enableMobileMode()
}, [isMobile])
```

## Split unrelated effects and memos

One hook with two jobs re-runs both when either input changes. Give each job its own hook and its own dependencies.

```tsx
useEffect(() => {
  analytics.trackPageView(pathname)
}, [pathname])

useEffect(() => {
  document.title = `${pageTitle} | My App`
}, [pageTitle])
```

## Subscribe to the derived value

If a component only needs `width < 768`, subscribe to that boolean (for example with `useMediaQuery`) and not to the raw width. It then re-renders when the answer changes, not on every pixel.

## Read state at the point of use

If a value is only read inside a handler, do not subscribe to it with a hook. Read it when the handler runs.

```tsx
function ShareButton({ chatId }: { chatId: string }) {
  function handleShare() {
    const ref = new URLSearchParams(window.location.search).get('ref')
    shareChat(chatId, { ref })
  }

  return <button onClick={handleShare}>Share</button>
}
```

## Use functional setState

When the next state depends on the previous one, pass an updater. The callback then needs no state dependency and cannot capture a stale value.

```tsx
const addItems = useCallback((newItems: Item[]) => {
  setItems((curr) => [...curr, ...newItems])
}, [])
```

Setting a value that does not depend on the old state (`setCount(0)`) needs no updater.

## Initialise state lazily

`useState(expensive())` runs `expensive()` on every render and discards the result. Pass a function.

```tsx
const [settings, setSettings] = useState(() => {
  const stored = localStorage.getItem('settings')
  return stored ? JSON.parse(stored) : {}
})
```

Cheap values such as `useState(0)` or `useState({})` do not need it.

## Keep transient values in refs

A value that changes often and does not affect the output, such as a pointer position or a timer id, belongs in `useRef`. Writing to a ref does not render. Update the DOM node directly when you need to show it.

```tsx
const dotRef = useRef<HTMLDivElement>(null)

useEffect(() => {
  const onMove = (e: MouseEvent) => {
    if (dotRef.current) dotRef.current.style.transform = `translateX(${e.clientX}px)`
  }
  window.addEventListener('mousemove', onMove)
  return () => window.removeEventListener('mousemove', onMove)
}, [])
```

## Mark non-urgent updates as transitions

Wrap frequent updates that the user does not need to see instantly in `startTransition`. React then keeps urgent input responsive.

```tsx
const onScroll = () => startTransition(() => setScrollY(window.scrollY))
```

## Defer expensive derived renders

When typing triggers a heavy filter or chart, pass the value through `useDeferredValue`. The input updates at once and the heavy part follows.

```tsx
const [query, setQuery] = useState('')
const deferredQuery = useDeferredValue(query)
const filtered = useMemo(
  () => items.filter((item) => fuzzyMatch(item, deferredQuery)),
  [items, deferredQuery],
)
const isStale = query !== deferredQuery
```

Without the `useMemo`, the filter still runs on every render.

## Memoisation rules, for use after profiling

- Move expensive work into its own `memo` component so a parent can return early before computing it. Do not compute it in the parent and guard it with `useMemo`.
- Give a memoised component's optional object, array or function props a default from a module constant. A default written inline in the parameter list is a new value each render and defeats `memo`.

  ```tsx
  const noop = () => {}
  const UserAvatar = memo(function UserAvatar({ onClick = noop }: { onClick?: () => void }) {
    // ...
  })
  ```

- Do not wrap a cheap expression with a primitive result in `useMemo`. The hook costs more than `a.loading || b.loading`.
