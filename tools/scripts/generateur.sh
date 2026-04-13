#!/usr/bin/env bash
# ==========================================
# Script Name: generate.sh
# Description: Générateur de scripts MySelf
# Author: MySelf System
# ==========================================

set -euo pipefail

TEMPLATE="TEMPLATE.sh"
TARGET_DIR="$(dirname "$0")"

usage() {
    echo "Usage: $0 <nom_du_script>"
    exit 1
}

log() {
    echo "[GENERATOR] $1"
}

error() {
    echo "[ERROR] $1" >&2
    exit 1
}

main() {
    [[ $# -eq 1 ]] || usage

    SCRIPT_NAME="$1"
    SCRIPT_PATH="$TARGET_DIR/${SCRIPT_NAME}.sh"

    [[ -f "$TARGET_DIR/$TEMPLATE" ]] || error "Template introuvable : $TEMPLATE"
    [[ -f "$SCRIPT_PATH" ]] && error "Le script existe déjà : $SCRIPT_PATH"

    log "Création du script : $SCRIPT_NAME"

    sed \
        -e "s/<nom_du_script>/$SCRIPT_NAME/g" \
        -e "s/<description_courte>/Script généré automatiquement/g" \
        "$TARGET_DIR/$TEMPLATE" > "$SCRIPT_PATH"

    chmod +x "$SCRIPT_PATH"

    log "Script créé : $SCRIPT_PATH"
}

main "$@"
