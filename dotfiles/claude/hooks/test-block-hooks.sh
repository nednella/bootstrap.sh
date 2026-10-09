#!/usr/bin/env bash
# Feeds each block-* hook the PreToolUse payload Claude Code sends and prints one line per case.
# Exits 1 if any case fails. BASH_RUNNER=/bin/bash reruns every case under macOS's bash 3.2.
set -u
hooks=$(dirname "${BASH_SOURCE[0]}")
failed=0

decision() { # <hook> <command>
  local output
  output=$(jq -cn --arg command "$2" '{hook_event_name: "PreToolUse", tool_name: "Bash", tool_input: {command: $command}}' | ${BASH_RUNNER:-} "$hooks/$1")
  [[ -z $output ]] && echo allow || jq -r '.hookSpecificOutput.permissionDecision' <<<"$output"
}

expect() { # <expected> <hook> <command>
  local actual
  actual=$(decision "$2" "$3")
  if [[ $actual == "$1" ]]; then echo "PASS $2: $3"; else echo "FAIL $2: $3 (expected $1, got $actual)"; failed=1; fi
}

expect deny  block-pr-merge.sh 'gh pr merge 12'
expect deny  block-pr-merge.sh 'gh pr merge 12 --squash --auto'
expect deny  block-pr-merge.sh 'gh -R a/b pr merge 12'
expect deny  block-pr-merge.sh '/opt/homebrew/bin/gh pr merge 12'
expect deny  block-pr-merge.sh 'GH PR MERGE 12'
expect deny  block-pr-merge.sh 'g"h" pr mer\ge 12'
expect deny  block-pr-merge.sh 'bash -c "gh pr merge 12"'
expect deny  block-pr-merge.sh 'echo 12 | xargs gh pr merge'
expect deny  block-pr-merge.sh 'gh api -X PUT repos/a/b/pulls/5/merge'
expect deny  block-pr-merge.sh 'gh api graphql -f query="mutation{enablePullRequestAutoMerge(input:{}){clientMutationId}}"'
expect allow block-pr-merge.sh 'gh pr merge --help'
expect allow block-pr-merge.sh 'git merge main'
expect allow block-pr-merge.sh 'gh pr list --state merged'
expect allow block-pr-merge.sh 'git commit -m "fix the pr merge copy"'

expect deny  block-pr-ready.sh 'gh pr ready 12'
expect deny  block-pr-ready.sh 'gh -R a/b pr ready 12'
expect deny  block-pr-ready.sh "gh api graphql -f query='mutation{markPullRequestReadyForReview(input:{}){clientMutationId}}'"
expect allow block-pr-ready.sh 'gh pr ready 12 --undo'
expect allow block-pr-ready.sh 'gh pr view 12'
expect allow block-pr-ready.sh 'git commit -m "the pr ready state"'

expect deny  block-review-requests.sh 'gh pr edit 12 --add-reviewer x'
expect deny  block-review-requests.sh 'gh pr edit --add-reviewer x 12'
expect deny  block-review-requests.sh 'gh pr edit 12 --add-reviewer=x'
expect deny  block-review-requests.sh 'gh pr create --draft --reviewer x'
expect deny  block-review-requests.sh 'gh pr create --draft -r x'
expect deny  block-review-requests.sh 'gh api repos/a/b/pulls/5/requested_reviewers -f "reviewers[]=x"'
expect deny  block-review-requests.sh 'gh api graphql -f query="mutation{requestReviews(input:{}){clientMutationId}}"'
expect allow block-review-requests.sh 'gh pr create --draft --title x --body y'
expect allow block-review-requests.sh 'gh pr edit 12 --body x'
expect allow block-review-requests.sh 'gh pr edit 12 --remove-reviewer x'
expect allow block-review-requests.sh 'ls -r'
expect allow block-review-requests.sh 'gh pr create -R a/b --draft --title x --body y'
expect deny  block-review-requests.sh 'gh pr create -R a/b --draft -r x'

body='Do not run gh pr merge, gh pr ready or gh pr edit --add-reviewer x; no --reviewer, no -r x.'
for hook in block-pr-merge.sh block-pr-ready.sh block-review-requests.sh; do
  expect allow $hook "gh pr create -R a/b --draft --title x --body \"$body\""
  expect allow $hook "gh pr create -R a/b --draft --title x --body '$body'"
  expect allow $hook "gh pr create -R a/b --draft --body-file - <<'EOF'"$'\n'"$body"$'\nEOF'
  expect allow $hook "gh pr create --draft --body \"\$(cat <<'EOF'"$'\n'"$body"$'\nEOF\n)"'
  expect allow $hook "agentos new x --prompt - <<'EOF'"$'\n'"$body"$'\nEOF'
  expect allow $hook "cat > /tmp/b.md <<-\"EOF\""$'\n'"$body"$'\n\tEOF'
done
expect deny  block-pr-merge.sh "cat > /tmp/b.md <<'EOF'"$'\nx\nEOF\ngh pr merge 1'
expect deny  block-pr-merge.sh $'gh pr create --body-file - <<EOF\n$(gh pr merge 1)\nEOF'
expect deny  block-pr-merge.sh 'gh pr create --body "$(gh pr merge 1)"'
expect deny  block-pr-merge.sh $'bash <<\'EOF\'\ngh pr merge 1\nEOF'
expect deny  block-pr-merge.sh $'/bin/sh -s <<\'EOF\'\ngh pr merge 1\nEOF'
expect deny  block-pr-merge.sh $'gh api graphql -F query=@- <<\'EOF\'\nmutation{mergePullRequest(input:{}){clientMutationId}}\nEOF'
expect deny  block-pr-merge.sh 'gh pr create --title x && gh pr merge 1'

for hook in block-pr-merge.sh block-pr-ready.sh block-review-requests.sh; do
  output=$(echo 'not json' | ${BASH_RUNNER:-} "$hooks/$hook" 2>/dev/null)
  [[ $(jq -r '.hookSpecificOutput.permissionDecision' <<<"$output" 2>/dev/null) == deny ]] && echo "PASS $hook: unreadable payload" || { echo "FAIL $hook: unreadable payload (expected deny)"; failed=1; }
done

exit $failed
