#!/usr/bin/env bash

# Note: be sure to add `.exe/` to your `.gitignore` file, otherwise this script will not work properly.

set -euo pipefail

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &> /dev/null && pwd)"

input_sha=$(
    (
        cd "$script_dir"
        git ls-tree -r HEAD -- .
        # shellcheck disable=SC2016
        git status --short -uall . | cut -c4- | xargs -n1 /bin/sh -c 'sha1sum "$0" 2>&1 || true'
    ) | sha1sum | cut -d' ' -f1
)

exe="$script_dir/.exe/$input_sha"

[[ -x $exe ]] || {
    echo "Compiling..."
    mkdir -p .exe
    go build -o "$exe" .
}

exec "$exe" "$@"
