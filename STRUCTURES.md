# Structure v1.1.0

MySelf/
├── cmd/
│   └── collector/
│       └── main.go
├── internal/
│   └── profile/
│       └── profile.go
├── records/
├── Makefile
└── go.mod

# Structure v1.2.0

# MySelf — structure du projet

## 1. Objectif du dépôt

MySelf est un système personnel modulaire pour se comprendre, se suivre et s’améliorer dans le temps.  
Ce dépôt définit une structure claire, stable et évolutive pour organiser :

- **l’identité** (valeurs, histoire, vision)
- **les états internes** (émotions, besoins, schémas)
- **les objectifs** (court, moyen, long terme)
- **les routines et protocoles** (habitudes, checklists, plans d’action)
- **les revues** (journal, bilans, rétrospectives)
- **les projets** (personnels, professionnels, relationnels)

Ce fichier décrit la **structure de référence** du projet MySelf.

---

## 2. Vue d’ensemble de l’arborescence

```text
MySelf/
├─ 0-meta/
│  ├─ README.md
│  ├─ structure.md
│  └─ changelog.md
├─ 1-identite/
│  ├─ vision/
│  ├─ valeurs/
│  ├─ histoire/
│  └─ roles/
├─ 2-etats-internes/
│  ├─ emotions/
│  ├─ besoins/
│  ├─ croyances/
│  └─ patterns/
├─ 3-objectifs/
│  ├─ court-terme/
│  ├─ moyen-terme/
│  └─ long-terme/
├─ 4-routines-protocoles/
│  ├─ routines-quotidiennes/
│  ├─ routines-hebdomadaires/
│  ├─ protocoles-crise/
│  └─ protocoles-developpement/
├─ 5-journal-revues/
│  ├─ journal-quotidien/
│  ├─ revues-hebdo/
│  ├─ revues-mensuelles/
│  └─ revues-annuelles/
├─ 6-projets/
│  ├─ perso/
│  ├─ pro/
│  └─ relations/
└─ 7-archives/
   ├─ anciens-objectifs/
   ├─ anciens-projets/
   └─ anciennes-versions/
