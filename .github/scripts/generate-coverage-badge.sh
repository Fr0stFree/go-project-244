#!/bin/sh

set -eu

coverage_profile="${1:-coverage.out}"
badge_path="${2:-.github/badges/coverage-badge.json}"

coverage=$(go tool cover -func="$coverage_profile" | awk '/^total:/ {print substr($3, 1, length($3)-1)}')
color=$(awk -v coverage="$coverage" 'BEGIN {
  if (coverage >= 80) print "brightgreen";
  else if (coverage >= 60) print "yellow";
  else print "red"
}')

mkdir -p "$(dirname "$badge_path")"

printf '{\n  "schemaVersion": 1,\n  "label": "coverage",\n  "message": "%s%%",\n  "color": "%s"\n}\n' "$coverage" "$color" > "$badge_path"

echo "Coverage badge generated at $badge_path with coverage $coverage% and color $color."