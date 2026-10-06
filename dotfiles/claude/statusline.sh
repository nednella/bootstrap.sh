#!/bin/bash
input=$(cat)
j() { jq -r "$1 // empty" <<<"$input"; }

c() { printf '\033[%sm%s\033[0m' "$1" "$2"; }
sep=$(c 90 ' | ')

dir=$(j .workspace.current_dir)
out=$(c 36 "${dir##*/}")

if branch=$(git -C "$dir" --no-optional-locks branch --show-current 2>/dev/null) && [ -n "$branch" ]; then
  out+="$sep$(c 33 "$branch")"
fi

model=$(j .model.display_name)
effort=$(j .effort.level)
out+="$sep$(c 32 "$model${effort:+ $effort}")"

pct=$(j .context_window.used_percentage)
if [ -n "$pct" ]; then
  used=$(j '.context_window.total_input_tokens')
  size=$(j '.context_window.context_window_size')
  fill=$(( (pct * 8 + 50) / 100 ))
  bar=$(printf '%*s' "$fill" '' | tr ' ' '#')$(printf '%*s' $((8 - fill)) '' | tr ' ' '-')
  if [ "$pct" -ge 80 ]; then color=31; elif [ "$pct" -ge 50 ]; then color=33; else color=32; fi
  [ "$size" -ge 1000000 ] && limit="$((size / 1000000)).$((size % 1000000 / 100000))M" || limit="$((size / 1000))k"
  out+="$sep$(c $color "ctx [$bar] $pct% $((used / 1000))k/$limit")"
fi

cost=$(j .cost.total_cost_usd)
ms=$(j .cost.total_duration_ms)
mins=$(( ${ms:-0} / 60000 ))
dur="${mins}m"; [ "$mins" -ge 60 ] && dur="$((mins / 60))h$((mins % 60))m"
out+="$sep$(printf '$%.2f %s' "${cost:-0}" "$dur")"

usage() {
  local used
  used=$(j ".rate_limits.$1.used_percentage")
  [ -n "$used" ] && out+="$sep$(printf '%s %.0f%%@%s' "$2" "$used" "$(date -r "$(j ".rate_limits.$1.resets_at")" "$3")")"
}
usage five_hour 5h '+%H:%M'
usage seven_day 7d '+%a %H:%M'

printf '%s' "$out"
