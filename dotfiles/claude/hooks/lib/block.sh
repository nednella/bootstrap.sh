#!/usr/bin/env bash
# Sourced by the block-* hooks. Reads the Bash command from the PreToolUse payload with line
# continuations joined and every quote and backslash dropped, so a word spelled in pieces
# (g"h", mer\ge) reads whole. A payload jq cannot read is denied, not let through.
#
# $command leaves out text the shell does not run: the body of a heredoc with a quoted
# delimiter (<<'EOF') unless a shell reads it, and the quoted value of a --body, --title,
# --message, -b, -t or -m flag that holds no $ or backtick. A PR body or a prompt can then
# name a blocked command. $text keeps everything, for the API names no prose needs.
shopt -s nocasematch

GH='(^|[^[:alnum:]_.-])gh([[:space:];|&<>)}`]|$)'

deny() { # <reason>
  jq -cn --arg reason "$1" '{hookSpecificOutput: {hookEventName: "PreToolUse", permissionDecision: "deny", permissionDecisionReason: $reason}}'
  exit 0
}

drop_quoted_heredocs() {
  awk '
    skip != "" { line = $0; if (dash) sub(/^\t+/, "", line); if (line == skip) skip = ""; next }
    { print }
    {
      i = index($0, "<<")
      if (!i || substr($0, i + 2, 1) == "<") next
      rest = substr($0, i + 2)
      dash = sub(/^-/, "", rest)
      sub(/^[ \t]*/, "", rest)
      if (!match(rest, /^("[A-Za-z0-9_]+"|\047[A-Za-z0-9_]+\047)/)) next
      if (substr($0, 1, i - 1) ~ /(^|[^A-Za-z0-9_.-])((ba|z|da|k)?sh|eval|source)([ \t;|&)]|$)/) next
      skip = substr(rest, 2, RLENGTH - 2)
    }'
}

read_command() { # <reason to deny with when the payload cannot be read>
  local raw clean flag_values
  raw=$(jq -r '.tool_input.command // empty') || deny "$1"
  clean='split("\\\n") | join("") | gsub("[\"'\''\\\\]"; "")'
  flag_values='gsub("(?<f>(^|\\s)(--body|--title|--message|-[bmt]))(=|\\s+)(\"[^\"\\\\$`]*\"|'\''[^'\'']*'\'')"; "\(.f) ")'
  text=$(jq -Rrs "$clean" <<<"$raw") || deny "$1"
  command=$(drop_quoted_heredocs <<<"$raw" | jq -Rrs "$flag_values | $clean") || deny "$1"
}
