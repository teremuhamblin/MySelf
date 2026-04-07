📄 1. docs/analysis.md (v1.2.0)

`md

Analyse — MySelf 1.2.0

Le module d’analyse fournit des statistiques, tendances et résumés automatiques basés sur les records collectés.  
Il constitue la première couche d’intelligence du système MySelf.

---

🎯 Objectifs
- Extraire des informations utiles à partir des données brutes.
- Identifier des patterns récurrents.
- Générer des résumés synthétiques.
- Préparer les données pour le dashboard CLI et futur dashboard web.

---

🧱 Architecture

```text
internal/analysis/
├── analyzer.go
├── stats.go
├── patterns.go
└── summary.go
```

---

📊 Fonctionnalités

1. Statistiques globales
- Nombre total de records
- Records par type
- Records par période (jour/semaine/mois)
- Tags les plus fréquents

2. Détection de patterns
- Répétitions
- Pics d’activité
- Périodes creuses
- Types dominants

3. Résumés automatiques
- Top 3 des thèmes du mois
- Évolution du volume d’entrées
- Analyse des tags dominants

---

🧪 Tests
- Tests unitaires sur les fonctions de stats
- Tests sur les patterns
- Tests sur les résumés

---

🔌 Intégration CLI
Le dashboard CLI utilise directement ce module pour afficher :
- histogrammes ASCII
- tableaux
- résumés textuels
`

---
