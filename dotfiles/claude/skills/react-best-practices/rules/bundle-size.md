# Bundle size

A smaller initial bundle means a faster first load.

## Import from the file, not the barrel

A barrel file re-exports many modules. Importing one name from a large barrel can pull thousands of modules into dev builds and slow HMR. Production tree-shaking often removes the excess, but not always.

```tsx
import Button from '@mui/material/Button'
import TextField from '@mui/material/TextField'
```

Check that the library ships types for its deep paths. Some, such as `lucide-react`, do not, and deep imports then become implicit `any` under strict mode. In that case keep the named import.

Libraries that commonly cause this: icon packs, `@mui/*`, `@headlessui/react`, `lodash`, `date-fns`, `rxjs`, `react-use`.

## Lazy-load heavy components

Load large components only when they render. Use `React.lazy` with a `Suspense` boundary.

```tsx
const MonacoEditor = lazy(() =>
  import('./monaco-editor').then((m) => ({ default: m.MonacoEditor })),
)

function CodePanel({ code }: { code: string }) {
  return (
    <Suspense fallback={<EditorSkeleton />}>
      <MonacoEditor value={code} />
    </Suspense>
  )
}
```

Split at routes first, then at heavy widgets such as editors, charts and PDF viewers. Define each `lazy` call at module level, never inside a component.

## Load data and modules when a feature turns on

```tsx
function AnimationPlayer({ enabled }: { enabled: boolean }) {
  const [frames, setFrames] = useState<Frame[] | null>(null)

  useEffect(() => {
    if (!enabled || frames) return
    let cancelled = false
    import('./animation-frames').then((mod) => {
      if (!cancelled) setFrames(mod.frames)
    })
    return () => {
      cancelled = true
    }
  }, [enabled, frames])

  if (!frames) return <Skeleton />
  return <Canvas frames={frames} />
}
```

## Defer non-critical third-party code

Analytics, logging and error reporting do not need to block first paint. Load them after the app mounts.

```tsx
const Analytics = lazy(() => import('./Analytics'))

function App() {
  return (
    <>
      <Routes />
      <Suspense fallback={null}>
        <Analytics />
      </Suspense>
    </>
  )
}
```

## Preload on intent

Start the download when the user shows intent, so the click feels instant.

```tsx
const loadEditor = () => import('./monaco-editor')

function EditorButton({ onClick }: { onClick: () => void }) {
  return (
    <button onMouseEnter={loadEditor} onFocus={loadEditor} onClick={onClick}>
      Open editor
    </button>
  )
}
```

Calling `import()` again for a loaded module is free. You can also preload from an effect when a feature flag turns on.

## Keep dynamic import paths analysable

A bundler cannot follow `import(variable)`. It either includes a broad set of files or warns. Use an explicit map of loaders.

```ts
const pageLoaders = {
  home: () => import('./pages/home'),
  settings: () => import('./pages/settings'),
} as const

const page = await pageLoaders[pageName]()
```
