# Data fetching and browser APIs

## Use a request cache library

Fetching in `useEffect` makes each component instance send its own request and gives you no caching or revalidation. Use a library such as SWR or TanStack Query. Instances with the same key share one request.

```tsx
function UserList() {
  const { data: users } = useSWR('/api/users', fetcher)
  // ...
}
```

Use the library's mutation helper for writes, and its immutable option for data that never changes during a session.

If you must fetch in an effect, abort on cleanup and ignore stale results.

```tsx
useEffect(() => {
  const controller = new AbortController()
  fetch('/api/users', { signal: controller.signal })
    .then((r) => r.json())
    .then(setUsers)
    .catch((e) => {
      if (e.name !== 'AbortError') setError(e)
    })
  return () => controller.abort()
}, [])
```

## Share one global listener

A hook that adds a `window` listener per instance adds N listeners for N uses. Keep one listener at module level and a registry of callbacks.

```tsx
const handlers = new Map<string, Set<() => void>>()
let detach: (() => void) | null = null

function onKeyDown(e: KeyboardEvent) {
  if (e.metaKey) handlers.get(e.key)?.forEach((cb) => cb())
}

function subscribe(key: string, cb: () => void) {
  if (!handlers.has(key)) handlers.set(key, new Set())
  handlers.get(key)!.add(cb)
  if (!detach) {
    window.addEventListener('keydown', onKeyDown)
    detach = () => window.removeEventListener('keydown', onKeyDown)
  }
  return () => {
    const set = handlers.get(key)!
    set.delete(cb)
    if (set.size === 0) handlers.delete(key)
    if (handlers.size === 0) {
      detach?.()
      detach = null
    }
  }
}

function useKeyboardShortcut(key: string, callback: () => void) {
  const callbackRef = useRef(callback)
  useEffect(() => {
    callbackRef.current = callback
  })
  useEffect(() => subscribe(key, () => callbackRef.current()), [key])
}
```

Do this only when many instances exist. For one or two, a plain effect is fine.

## Use passive listeners for touch and wheel

A listener that never calls `preventDefault()` should say so. The browser can then scroll without waiting for it.

```ts
useEffect(() => {
  const onWheel = (e: WheelEvent) => track(e.deltaY)
  document.addEventListener('wheel', onWheel, { passive: true })
  return () => document.removeEventListener('wheel', onWheel)
}, [])
```

Do not use passive for custom swipe or zoom handlers that call `preventDefault()`.

## Version and trim localStorage data

Put a version in each key, store only the fields the UI needs, and wrap every access in `try/catch`. Storage throws in private mode, when full, or when disabled.

```ts
const KEY = 'userConfig:v2'

function saveConfig(config: { theme: string; language: string }) {
  try {
    localStorage.setItem(KEY, JSON.stringify(config))
  } catch {
    return
  }
}

function loadConfig() {
  try {
    const raw = localStorage.getItem(KEY)
    return raw ? JSON.parse(raw) : null
  } catch {
    return null
  }
}
```

Never store tokens or personal data. When the shape changes, bump the version and migrate or drop the old key.
