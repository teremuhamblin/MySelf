# Changelog — MySelf

Toutes les modifications notables du projet sont documentées dans ce fichier.  
Format basé sur les recommandations de [Keep a Changelog](https://keepachangelog.com/) et le versionnement sémantique (SemVer).

---

## [1.2.0] — Analyse & Dashboard — 2026-04-07
### Ajouté
- Module complet d’analyse (`internal/analysis/`) :
  - statistiques globales (types, tags, fréquence, volume)
  - détection de patterns simples
  - génération de résumés automatiques
- Dashboard CLI (`cmd/dashboard/`) :
  - vue synthétique des données
  - histogrammes ASCII
  - filtres par type, période, tags
- API interne d’accès aux données (`pkg/data/`)
  - fonctions de filtrage, listing, récupération par ID
- Nouvelle commande CLI :
  - `myself dashboard`
  - `myself stats`
  - `myself trends`
  - `myself export --csv`
- Documentation :
  - `docs/dashboard.md`
  - `docs/analysis.md`

### Modifié
- Réorganisation du dossier `records/` :
  - structure par année/mois
  - index enrichi (`records/index.json`)
- Amélioration du modèle `profile` :
  - ajout de `Type`, `Tags`, `CreatedAt`, `UpdatedAt`
- Makefile enrichi (test, lint, fmt)

### Corrigé
- Validation des records plus robuste
- Correction de chemins relatifs dans le collector

---

## [1.1.0] — Modularisation & Structure — 2026-04-05
### Ajouté
- Nouvelle architecture modulaire :
  - `internal/collector/`
  - `pkg/` pour modules réutilisables
- CLI améliorée :
  - sous-commandes (`add`, `list`, `view`)
  - flags (`--type`, `--tag`, `--json`)
- Indexation des records (`records/index.json`)
- Tests unitaires initiaux
- Ajout de `golangci-lint`

### Modifié
- Refonte du modèle `profile`
- Makefile enrichi
- Documentation mise à jour (`README`, `structure.md`)

### Corrigé
- Gestion des doublons dans les records
- Sérialisation JSON plus propre

---

## [1.0.0] — Version initiale — 2026-04-03
### Ajouté
- Structure minimale du projet :
  - `cmd/collector/main.go`
  - `internal/profile/profile.go`
  - `records/`
  - `Makefile`
  - `go.mod`
- Collecte simple de records via CLI
- Stockage brut des données
- Documentation initiale

---

## Format des versions
- **MAJEUR** : changements incompatibles
- **MINEUR** : nouvelles fonctionnalités rétro-compatibles
- **PATCH** : corrections et améliorations mineures
