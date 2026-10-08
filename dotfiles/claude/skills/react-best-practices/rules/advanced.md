# Advanced patterns

## Keep a subscription stable while the callback changes

If an effect subscribes with a callback prop, putting the callback in the dependency array re-subscribes on each render. Store the latest callback in a ref and subscribe once.

```tsx
function useWindowEvent(event: string, handler: (e: Event) => void) {
  const handlerRef = useRef(handler)
  useEffect(() => {
    handlerRef.current = handler
  })

  useEffect(() => {
    const listener = (e: Event) => handlerRef.current(e)
    window.addEventListener(event, listener)
    return () => window.removeEventListener(event, listener)
  }, [event])
}
```

The same pattern fixes timers that call a callback prop, such as a debounced search. Keep the callback in a ref and leave it out of the dependencies.

`useEffectEvent` does the same job but needs React 19.2. Use a ref on React 18.

## Run app setup once, outside components

Code that must run once per page load does not belong in `useEffect(..., [])`. Effects run twice in development under Strict Mode and again whenever the component remounts. Put the setup in the entry module (`main.tsx`) before `createRoot`.

```tsx
loadFromStorage()
checkAuthToken()

createRoot(document.getElementById('root')!).render(<App />)
```

If it must wait for a component, use a module-level `didInit` flag.
