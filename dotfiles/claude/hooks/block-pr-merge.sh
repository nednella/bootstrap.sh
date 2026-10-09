#!/usr/bin/env bash
# Blocks every way to merge a PR: gh pr merge, the REST merge endpoint and the GraphQL merge,
# auto-merge and merge-queue mutations. `gh pr merge --help` on its own still runs.
source "$(dirname "${BASH_SOURCE[0]}")/lib/block.sh"

reason="Merging a PR is the user's alone. Leave the PR open and tell the user it is ready to merge."
read_command "$reason"

[[ $command =~ ^[[:space:]]*gh[[:space:]]+pr[[:space:]]+merge[[:space:]]+--help[[:space:]]*$ ]] && exit 0
[[ $command =~ $GH && $command =~ (^|[^[:alnum:]_-])pr[[:space:]]+merge([^[:alnum:]_-]|$) ]] && deny "$reason"
[[ $text =~ /pulls/[0-9]+/merge ]] && deny "$reason"
[[ $text =~ mergePullRequest|enablePullRequestAutoMerge|enqueuePullRequest ]] && deny "$reason"
exit 0
