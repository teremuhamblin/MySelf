📄 2. docs/dashboard-cli.md (v1.2.0)

`md

Dashboard CLI — MySelf 1.2.0

Le dashboard CLI permet de visualiser rapidement les données collectées via une interface en ligne de commande.

---

🎯 Objectifs
- Offrir une vue synthétique des données.
- Permettre une navigation simple.
- Fournir des insights rapides.

---

🧱 Architecture

```text
cmd/dashboard/
└── main.go
```

---

🧰 Commandes disponibles

myself dashboard
Affiche la vue globale :
- nombre total d’entrées
- histogramme par type
- tags les plus fréquents
- timeline simple

myself dashboard --month YYYY-MM
Vue mensuelle détaillée.

myself dashboard --type <type>
Filtre par type :
- emotion
- habit
- event
- note

---

📊 Exemples d’affichage

Histogramme ASCII
`
emotions   ████████████  42
habits     ████          12
events     ██████        18
notes      █████████     30
`

Résumé automatique
`
Mois d’avril : 102 entrées
Thèmes dominants : stress, énergie, apprentissage
Journées les plus actives : 3, 12, 21
`

---

🔌 Intégration
Le dashboard CLI utilise :
- pkg/data/ pour accéder aux records
- internal/analysis/ pour les statistiques
`

---
