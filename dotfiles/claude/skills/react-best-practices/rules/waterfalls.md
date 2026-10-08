# Waterfalls

Each awaited call that could have run in parallel adds a full network round trip. Fix these first.

## Run independent calls together

Start calls that do not depend on each other at the same time, then await them as a group.

```ts
const [user, posts, comments] = await Promise.all([
  fetchUser(),
  fetchPosts(),
  fetchComments(),
])
```

## Start early, await late for partial dependencies

When only some calls depend on others, create all promises first and chain only the dependent one.

```ts
const userPromise = fetchUser()
const profilePromise = userPromise.then((user) => fetchProfile(user.id))
const [user, config, profile] = await Promise.all([
  userPromise,
  fetchConfig(),
  profilePromise,
])
```

The `better-all` package expresses the same graph declaratively. Plain promises are enough for most code.

## Await only where the result is used

Move an `await` into the branch that needs it. An early exit then skips the wait.

```ts
async function save(resourceId: string, userId: string) {
  const resource = await getResource(resourceId)
  if (!resource) return { error: 'Not found' }

  const permissions = await fetchPermissions(userId)
  if (!permissions.canEdit) return { error: 'Forbidden' }

  return updateResource(resource, permissions)
}
```

## Check cheap conditions before async ones

In `flag && cheapCondition`, test the cheap synchronous condition first. Otherwise you pay for the async call when the result cannot matter.

```ts
if (isEditable && (await getFlag('inline-edit'))) {
  // ...
}
```

## Put Suspense boundaries around the slow part

Do not block a whole page on one slow section. Render the layout at once and wrap only the section that waits.

```tsx
const Chart = lazy(() => import('./Chart'))

function Dashboard() {
  return (
    <div>
      <Sidebar />
      <Header />
      <Suspense fallback={<ChartSkeleton />}>
        <Chart />
      </Suspense>
      <Footer />
    </div>
  )
}
```

The same applies to data, if your data library supports Suspense. Share one request between sibling components inside a boundary so they wait together.

Skip a narrow boundary when the content decides the layout, or when the fallback would cause a visible jump. A small, fast section does not need one.
