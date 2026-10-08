#!/usr/bin/env bash
# Blocks requesting reviewers on a PR: the gh pr edit and gh pr create flags, the REST
# requested_reviewers endpoint and the GraphQL mutation.
source "$(dirname "${BASH_SOURCE[0]}")/lib/block.sh"

reason="Requesting reviewers is the user's alone. Tell the user the PR is ready for review."
read_command "$reason"

[[ $command =~ --add-reviewer|--reviewer([^[:alnum:]_-]|$) ]] && deny "$reason"
[[ $command =~ $GH && $command =~ pr[[:space:]]+create && $command =~ [[:space:]]-r([[:space:]=]|$) ]] && deny "$reason"
[[ $command =~ requested_reviewers|requestReviews ]] && deny "$reason"
exit 0
