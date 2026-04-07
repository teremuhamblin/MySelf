📄 3. docs/api.md (v1.3.0)

`md

API REST — MySelf 1.3.0

L’API REST permet d’accéder aux données, statistiques et analyses de MySelf depuis n’importe quelle interface.

---

🎯 Objectifs
- Rendre MySelf accessible via HTTP.
- Permettre l’intégration avec le dashboard web.
- Offrir une interface stable et documentée.

---

🧱 Architecture

`
cmd/api/
└── main.go

internal/api/
├── router.go
├── handlers.go
└── middleware.go
`

---

🔌 Endpoints

📁 Records

GET /records
Liste tous les records.

GET /records/{id}
Retourne un record spécifique.

---

📊 Statistiques

GET /stats
Statistiques globales :
- total
- par type
- par période
- tags dominants

---

📈 Tendances

GET /trends
Tendances temporelles :
- évolution mensuelle
- pics d’activité

---

🏷️ Tags

GET /tags
Liste des tags + fréquence.

---

🧠 Analyse

GET /analysis/summary
Résumé automatique.

GET /analysis/patterns
Patterns détectés.

GET /analysis/monthly
Analyse mensuelle.

---

🔧 Configuration
Fichier config.yaml :
`yaml
port: 8080
mode: debug
cors: true
`

---

🧪 Tests
- Tests handlers
- Tests router
- Tests intégration API
`

---

📄 4. docs/dashboard-web.md (v1.3.0)

`md

Dashboard Web — MySelf 1.3.0

Le dashboard web est l’interface visuelle principale de MySelf.  
Il permet de consulter les données, statistiques et analyses via une interface moderne.

---

🎯 Objectifs
- Visualiser les données de manière claire.
- Naviguer facilement entre les vues.
- Offrir une expérience utilisateur simple et intuitive.

---

🧱 Architecture

`
web/dashboard/
├── index.html
├── assets/
│   ├── styles.css
│   └── charts.js
└── js/
    ├── api.js
    ├── dashboard.js
    └── charts.js
`

---

🖥️ Pages principales

1. Vue globale
- total des records
- histogramme par type
- tags dominants
- timeline

2. Vue détaillée
- liste des records
- filtres :
  - type
  - tags
  - période
- affichage d’un record

3. Vue analyse
- résumés automatiques
- patterns détectés
- évolution mensuelle

---

🔌 Communication API
Le dashboard utilise fetch() pour interroger l’API REST :
`js
fetch("http://localhost:8080/stats")
  .then(r => r.json())
  .then(data => renderStats(data))
`

---

🎨 Style & UX
- CSS minimaliste
- Charts simples (ASCII ou canvas)
- Navigation fluide

---

🧪 Tests
- Lint JS
- Vérification du build
`

---
