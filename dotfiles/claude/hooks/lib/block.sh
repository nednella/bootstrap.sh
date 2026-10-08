#!/usr/bin/env bash
# Sourced by the block-* hooks. Reads the Bash command from the PreToolUse payload with line
# continuations joined and every quote and backslash dropped, so a word spelled in pieces
# (g"h", mer\ge) reads whole. A payload jq cannot read is denied, not let through.
shopt -s nocasematch

GH='(^|[^[:alnum:]_.-])gh([[:space:];|&<>)}`]|$)'

deny() { # <reason>
  jq -cn --arg reason "$1" '{hookSpecificOutput: {hookEventName: "PreToolUse", permissionDecision: "deny", permissionDecisionReason: $reason}}'
  exit 0
}

read_command() { # <reason to deny with when the payload cannot be read>
  command=$(jq -r '.tool_input.command // empty | split("\\\n") | join("") | gsub("[\"'\''\\\\]"; "")') || deny "$1"
}
