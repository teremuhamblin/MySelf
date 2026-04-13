#!/usr/bin/env bash
# ==========================================
# Script Name: rotate-logs
# Description: Rotation simple des logs MySelf
# Author: MySelf System
# ==========================================

set -euo pipefail

TARGET_DIR="$(dirname "$0")/../../tracking"

log() {
    echo "[ROTATE] $1"
}

main() {
    DATE="$(date +%Y-%m-%d)"

    for FILE in "$TARGET_DIR"/*.log; do
        [[ -f "$FILE" ]] || continue

        BASENAME="$(basename "$FILE")"
        ARCHIVE="${FILE}.${DATE}"

        cp "$FILE" "$ARCHIVE"
        : > "$FILE"

        log "Rotation effectuée : $BASENAME → $BASENAME.$DATE"
    done
}

main "$@"
