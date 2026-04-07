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

---

🛣️ Roadmap MySelf — Version 1.2.0 (Analyse + Dashboard)

`md

MySelf — Roadmap 1.2.0

Analyse + Dashboard

🎯 Vision générale
La version 1.2.0 transforme MySelf d’un simple collecteur de données (1.0.0) et d’un système structuré (1.1.0) en un outil capable de comprendre, analyser et visualiser les informations personnelles collectées.

Objectifs :
- produire des analyses automatiques,
- offrir une vue synthétique de soi,
- permettre une navigation simple dans les données,
- poser les bases d’un futur tableau de bord avancé (1.3+).

---

🚀 Nouveautés principales de la version 1.2.0

1. Module d’analyse (internal/analysis/)
Création d’un module dédié aux analyses automatiques.

Fonctionnalités prévues
- Statistiques globales :
  - nombre total de records,
  - records par type,
  - records par période (jour/semaine/mois),
  - tags les plus utilisés.
- Détection de patterns simples :
  - répétitions,
  - pics d’activité,
  - périodes creuses.
- Génération d’un résumé automatique :
  - “Top 3 des thèmes du mois”
  - “Évolution du nombre d’entrées”
  - “Types les plus fréquents”

Structure proposée

```text
`
internal/analysis/
├── analyzer.go
├── stats.go
├── patterns.go
└── summary.go
`
```

---

2. Dashboard CLI (cmd/dashboard/)
Un tableau de bord en ligne de commande, simple mais puissant.

Fonctionnalités
- myself dashboard → vue synthétique
- myself dashboard --month 2026-04 → vue mensuelle
- myself dashboard --type emotion → vue filtrée
- Affichage :
  - histogrammes ASCII,
  - tableaux,
  - résumés textuels.

Exemples d’affichage
- Histogramme des types :

``` schema
  `
  emotions   ████████████  42
  habits     ████          12
  events     ██████        18
  notes      █████████     30
  `
```
- Résumé :
  `
  Mois d’avril : 102 entrées
  Thèmes dominants : stress, énergie, apprentissage
  Journées les plus actives : 3, 12, 21
  `

---

3. Réorganisation du stockage (records/)
Pour permettre l’analyse efficace.

Modifications
- Structure par année/mois :

``` text
  `
  records/
  ├── 2026/
  │   ├── 04/
  │   │   ├── 2026-04-01.json
  │   │   ├── 2026-04-02.json
  │   │   └── ...
  `
```

- Ajout d’un fichier records/index.json enrichi :
  - ID
  - type
  - tags
  - date
  - chemin du fichier

---

4. API interne d’accès aux données (pkg/data/)
Pour séparer la logique d’accès aux records.

Fonctions clés
- ListRecords()
- GetRecordByID()
- FilterByType()
- FilterByDateRange()
- GetStats()

---

5. Améliorations de la CLI collector
Ajout de nouvelles commandes utiles pour l’analyse.

Nouvelles commandes
- myself stats
- myself tags
- myself trends
- myself export --csv

---

6. Qualité & outils
- Tests unitaires pour le module d’analyse.
- Benchmarks simples (analyse sur 1000 records).
- Ajout de make dashboard pour lancer le dashboard.
- Documentation mise à jour :
  - README
  - structure.md
  - docs/dashboard.md
  - docs/analysis.md

---

📌 Résumé des modifications entre 1.1.0 et 1.2.0

| Domaine | 1.1.0 | 1.2.0 |
|--------|--------|--------|
| Collecte | CLI modulaire | CLI enrichie (stats, trends) |
| Stockage | index + structure | structure par date + index enrichi |
| Analyse | aucune | module complet d’analyse |
| Dashboard | aucun | dashboard CLI |
| Architecture | modulaire | ajout analysis + data API |
| Documentation | complète | documentation avancée |

---

🏁 Livrables finaux pour 1.2.0
- Nouveau module analysis/
- Nouveau dashboard CLI
- Nouvelle structure records/
- API interne pkg/data/
- Documentation complète
- Tests + benchmarks
`

---
