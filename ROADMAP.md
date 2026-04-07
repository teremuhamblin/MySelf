# MySelf — Roadmap 1.0.0 → 1.1.0

## 🎯 Vision générale
MySelf évolue d’un simple collecteur de données personnelles (v1.0.0) vers un système plus structuré, modulaire et exploitable (v1.1.0), capable de stocker, manipuler et analyser les informations collectées.

---

# ✅ Version 1.0.0 — Base fonctionnelle minimale

## Objectif
Fournir un socle stable permettant :
- de collecter des données personnelles,
- de les stocker sous forme de *records*,
- de définir une structure interne de profil,
- de disposer d’un binaire fonctionnel.

## Contenu de la version 1.0.0

### 1. Architecture minimale
- `cmd/collector/main.go` : CLI simple permettant d’ajouter un record.
- `internal/profile/profile.go` : structure de base du profil (nom, date, tags, contenu).
- `records/` : stockage brut des entrées.

### 2. Fonctionnalités
- Collecte d’un record via CLI.
- Enregistrement dans `records/` (format JSON ou texte).
- Validation minimale des entrées.
- Makefile : build, run, clean.
- go.mod : dépendances propres.

### 3. Qualité & maintenance
- Structure stable.
- Code lisible.
- Documentation minimale.

---

# 🚀 Version 1.1.0 — Amélioration structurelle & modularité

## Objectif
Étendre MySelf pour le rendre :
- plus modulaire,
- plus structuré,
- plus facile à faire évoluer,
- plus utile pour l’utilisateur.

## Nouveautés prévues

### 1. Amélioration de l’architecture
- Création d’un dossier `internal/collector/` pour séparer la logique du CLI.
- Introduction d’un dossier `pkg/` pour les modules réutilisables.
- Refonte du dossier `records/` :
  - sous-dossiers par année/mois,
  - index automatique,
  - métadonnées.

### 2. Évolution du modèle `profile`
- Ajout de champs :
  - `Type` (emotion, habit, event, note…)
  - `Tags` (liste)
  - `CreatedAt`, `UpdatedAt`
- Validation renforcée.
- Sérialisation/désérialisation propre.

### 3. Nouvelle CLI (collector v2)
- Commandes :
  - `myself add` → ajouter un record
  - `myself list` → lister les records
  - `myself view <id>` → afficher un record
- Flags :
  - `--type`
  - `--tag`
  - `--json`

### 4. Gestion avancée des records
- Génération automatique d’ID.
- Indexation dans un fichier `records/index.json`.
- Support du format JSON structuré.
- Détection des doublons.

### 5. Qualité & outils
- Ajout de tests unitaires (min 40% coverage).
- Ajout de `golangci-lint`.
- Makefile enrichi :
  - `make test`
  - `make lint`
  - `make fmt`
- Documentation mise à jour :
  - README
  - structure.md
  - CONTRIBUTING.md (optionnel)

### 6. Améliorations possibles (optionnelles mais recommandées)
- Export des records en CSV.
- Import depuis un fichier externe.
- Ajout d’un module `analysis/` pour statistiques simples :
  - nombre d’entrées par type,
  - fréquence des tags,
  - timeline.

---

# 📌 Résumé des modifications entre 1.0.0 et 1.1.0

| Domaine | 1.0.0 | 1.1.0 |
|--------|-------|--------|
| Collecte | CLI simple | CLI modulaire avec sous-commandes |
| Stockage | fichiers bruts | indexation + structure par date |
| Profil | structure minimale | modèle enrichi + validation |
| Architecture | basique | modulaire (collector, pkg, analysis) |
| Qualité | minimale | tests + lint + outils |
| Documentation | basique | complète et structurée |

---

# 🏁 Livrables finaux pour 1.1.0
- Nouveau binaire `myself`.
- Nouvelle structure de dossiers.
- Documentation mise à jour.
- Tests + linting.
- Index des records.
- CLI complète.
