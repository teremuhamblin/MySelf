#!/usr/bin/env bash
# ==========================================
# Script Name: add-entry
# Description: Ajouter une entrée dans un log MySelf
# Author: MySelf System
# ==========================================

set -euo pipefail

usage() {
    echo "Usage: $0 <fichier.log> <texte de l'entrée>"
    exit 1
}

log() {
    echo "[ADD-ENTRY] $1"
}

error() {
    echo "[ERROR] $1" >&2
    exit 1
}

main() {
    [[ $# -ge 2 ]] || usage

    FILE="$1"
    shift
    ENTRY="$*"

    [[ -f "$FILE" ]] || error "Fichier introuvable : $FILE"

    DATE="$(date +%Y-%m-%d)"
    echo "$DATE — $ENTRY" >> "$FILE"

    log "Entrée ajoutée dans $FILE"
}

main "$@"
