# Structure du projet MySelf

## STRUCTURES v1.1.0

### 1. Objectif du projet MySelf

L’objectif de MySelf est de créer un système personnel capable de collecter, structurer et analyser des informations sur soi-même (habitudes, émotions, comportements, décisions, données quotidiennes, etc.) afin de :

- mieux se comprendre,  
- suivre son évolution dans le temps,  
- identifier des patterns,  
- prendre de meilleures décisions,  
- construire une version plus cohérente et intentionnelle de soi.

Le projet repose sur une architecture Go simple, modulaire et extensible, permettant d’ajouter facilement de nouveaux collecteurs, formats de données ou modules d’analyse.

---

## 2. Courte description du dépôt

Ce dépôt contient la base du système MySelf, organisé autour de trois éléments principaux :

- cmd/collector/ : l’application principale qui collecte les données (journal, habitudes, états internes, événements, etc.) et les enregistre sous forme de records.
- internal/profile/ : la logique interne représentant le profil personnel, incluant les structures, modèles et règles permettant de manipuler les données collectées.
- records/ : le stockage brut des données collectées, sous forme de fichiers ou d’entrées structurées.

Le projet utilise un Makefile pour automatiser les tâches (build, test, formatage) et un go.mod pour gérer les dépendances Go.

---

### 3. Vue d’ensemble de l’arborescence

```text
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
```

---

# STRUCTURES 1.2.0

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

### 2. Vue d’ensemble de l’arborescence

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
```

---

3. Détails par dossier

3.1 0-meta/

- README.md : présentation globale de MySelf.
- structure.md : ce fichier, référence de la structure.
- changelog.md : journal des évolutions de la structure et des concepts.

3.2 1-identite/

- vision/ : fichiers décrivant la vision long terme, les grandes directions de vie.
- valeurs/ : valeurs fondamentales, principes non négociables, priorités.
- histoire/ : récits, lignes de vie, événements marquants.
- roles/ : rôles actuels (ex : ami, professionnel, créateur, etc.) et attentes associées.

3.3 2-etats-internes/

- emotions/ : cartes émotionnelles, lexique personnel, triggers fréquents.
- besoins/ : besoins récurrents, non satisfaits, stratégies d’y répondre.
- croyances/ : croyances aidantes/limitantes, reformulations.
- patterns/ : schémas répétitifs, réactions automatiques, scripts internes.

3.4 3-objectifs/

- court-terme/ : objectifs sur 1–3 mois, très concrets.
- moyen-terme/ : objectifs sur 6–18 mois.
- long-terme/ : grandes directions sur plusieurs années.

Chaque objectif peut suivre un format standard (ex : YYYY-MM-nom-objectif.md).

3.5 4-routines-protocoles/

- routines-quotidiennes/ : matin, soir, énergie, hygiène mentale.
- routines-hebdomadaires/ : planification, revue, ménage mental.
- protocoles-crise/ : quoi faire quand ça va vraiment mal (plans simples, concrets).
- protocoles-developpement/ : protocoles pour travailler un point précis (confiance, assertivité, etc.).

3.6 5-journal-revues/

- journal-quotidien/ : entrées datées, format libre ou structuré.
- revues-hebdo/ : ce qui a marché, ce qui bloque, ajustements.
- revues-mensuelles/ : tendances, apprentissages, décisions.
- revues-annuelles/ : bilan global, réorientation éventuelle.

3.7 6-projets/

- perso/ : projets personnels (santé, apprentissage, créativité, etc.).
- pro/ : projets professionnels, carrière, compétences.
- relations/ : projets liés aux relations importantes, dynamiques à travailler.

Chaque projet peut suivre un format standard (ex : PROJET-nom.md avec : contexte, objectifs, étapes, risques, next actions).

3.8 7-archives/

- anciens-objectifs/ : objectifs terminés ou abandonnés.
- anciens-projets/ : projets clos.
- anciennes-versions/ : anciennes structures, anciens systèmes, pour garder une trace.

---

4. Conventions de nommage

- Dossiers : kebab-case (routines-quotidiennes, revues-mensuelles).
- Fichiers datés : YYYY-MM-DD-titre-court.md.
- Fichiers objectifs : YYYY-MM-objectif-titre.md.
- Fichiers projets : PROJET-nom-court.md.

---

5. Évolution de la structure

La structure de MySelf est vivante.  
Toute modification importante doit être :

1. Documentée dans 0-meta/changelog.md.
2. Répercutée dans ce fichier structure.md.
3. Motivée (pourquoi ce changement ? qu’est-ce qu’il simplifie ou améliore ?).

Quand une partie devient obsolète, elle est déplacée dans 7-archives/ plutôt que supprimée.

---
`
