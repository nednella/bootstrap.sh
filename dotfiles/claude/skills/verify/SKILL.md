---
name: verify
description: Check a finished change in the running app against what was asked, and collect evidence. Use after implementing a UI or behaviour change and before saying it works, or when asked to "verify", "check it works" or "prove it".
---

# Verify

Prove that a change does what was asked, in the running app, with evidence. Tests passing is not enough for a behaviour change. A user must be able to reach it and see it work.

## 1. Write the checks first

Before you open the app, turn the request into checks. Each check is one thing a user can do and one result they should see. Include the paths the change can affect:

- success
- cancel or back out
- error, such as a failed request or invalid input
- empty, with no data
- persistence, after a reload

Show the list. If the request does not say what one of them should do, ask.

## 2. Find the feature map

Look for `.claude/verify/README.md` in the repository. It indexes one short file per feature area, and says how to start the app. Each feature file answers four questions:

- What exists in this area.
- How a user reaches it.
- How to drive it.
- What usually goes wrong, or looks right when it is not.

Read only the files for the areas the change touches. If there is no map, or no file for the area, work out the same four answers from the code and the project's `CLAUDE.md`.

## 3. Start the app

Use the start command from the map or the project's `CLAUDE.md`. Some projects never run locally; use the URL they name. Make sure the app runs your change, not a stale build: check a value or a string you changed.

If you cannot run the app, stop. Say so, and say which checks you could not make. Do not claim them as passed.

## 4. Drive every check

In an agentos session, use your browser: run `agentos browser help` for the rules, then drive the pages with the `browser` MCP tools. Outside agentos, use the browser tools you have. For a change with no UI, drive it the way a caller would: the CLI, an API request, a job run.

For each check:

1. Reach the feature the way a user would, not by URL alone when the user would click there.
2. Act, then read the result from the page: the text, the state, the network response, the console.
3. Capture evidence. In agentos, `agentos browser screenshot --caption "<check>: <result>"`. Use `agentos show <file>` for other evidence.

## 5. Judge and report

Compare each result with its check, not with what the code was meant to do. Report one line per check:

| Check | Result | Evidence |
|---|---|---|

A check passes only with evidence. If a result is wrong, report it as a failure with what you saw. Fix it only if the user asks or the task already covers it, then verify again.

If the change broke something the map lists nearby, report that too.

## 6. Keep the map current

Update the map as part of the change, in the same commit:

- If the area has no feature file, write one with the four answers you used.
- If a file was wrong, such as a moved entry point or a renamed button, fix it.
- If the change added or removed something a user can reach, add or remove it.
- Keep `.claude/verify/README.md` listing every feature file, with the start command at the top.

Keep each file short and about behaviour, so it stays true when the code under it changes. Do not copy code into it.

Close the browser with `agentos browser close` when you finish.

Pattern adapted from poteto/verification-skill-example.
