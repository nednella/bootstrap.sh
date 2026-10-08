#!/usr/bin/env bash
# Blocks marking a draft PR ready for review: gh pr ready and the GraphQL mutation.
# `gh pr ready --undo`, which turns a PR back into a draft, still runs.
source "$(dirname "${BASH_SOURCE[0]}")/lib/block.sh"

reason="Marking a PR ready for review is the user's alone. Leave it as a draft and tell the user."
read_command "$reason"

[[ $command =~ --undo([^[:alnum:]_-]|$) ]] && exit 0
[[ $command =~ $GH && $command =~ (^|[^[:alnum:]_-])pr[[:space:]]+ready([^[:alnum:]_-]|$) ]] && deny "$reason"
[[ $command =~ markPullRequestReadyForReview ]] && deny "$reason"
exit 0
