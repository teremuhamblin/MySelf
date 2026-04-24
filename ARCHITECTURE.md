##### ARCHITECTURE.md - markdown
# 🧱 Architecture — MySelf

>Structure simple et modulaire du projet MySelf.

```text
MySelf/
 ├── core/          → Identité, doctrine, règles
 ├── goals/         → Objectifs CT/MT/LT
 ├── routines/      → Routines essentielles
 ├── tracking/      → Suivi discipline, santé, progression
 ├── .github/       → Templates, PR, CODEOWNERS
 ├── README.md      → Présentation du projet
 ├── ROADMAP.md     → Plan d’évolution
 ├── CHANGELOG.md   → Historique des versions
 ├── CONTRIBUTING.md→ Guide de contribution
 └── SECURITY.md    → Politique de sécurité
```

# 🧱 Architecture mise à jour v1.1

>Principes
- Simplicité
- Modularité
- Clarté
- Évolutivité

```text
MySelf/
│
├── core/
│   ├── identity.rst
│   ├── doctrine.rst
│   ├── rules.rst
│   ├── mindset.rst
│   ├── energy-management.rst
│   ├── focus-system.rst
│   └── life-areas.rst
│
├── goals/
│   ├── short-term.rst
│   ├── mid-term.rst
│   └── long-term.rst
│
├── routines/
│   ├── morning.rst
│   ├── evening.rst
│   ├── discipline.rst
│   ├── health.rst
│   └── reset-weekly.rst
│
├── tracking/
│   ├── daily-log.rst
│   ├── weekly-review.rst
│   ├── monthly-review.rst
│   ├── metrics.rst
│   └── habits.rst
│
├── tools/
│   ├── generate-daily-log.rst
│   ├── generate-weekly-review.rst
│   ├── checklist.rst
│   ├── scripts/
│   │   ├── generate_daily.py
│   │   ├── generate_weekly.py
│   │   └── generate_monthly.py
│   └── templates/
│       ├── daily_template.rst
│       ├── weekly_template.rst
│       └── monthly_template.rst
│
├── docs/
│   ├── index.rst
│   ├── conf.py
│   ├── Makefile
│   └── _static/
│       └── theme.css
│
├── dashboard.rst
├── STYLEGUIDE.rst
├── WORKFLOW.rst
├── ROADMAP.rst
├── ARCHITECTURE.rst
│
└── .github/
    ├── ISSUE_TEMPLATE/
    │   ├── bug_report.yml
    │   ├── feature_request.yml
    │   ├── improvement.yml
    │   └── question.yml
    ├── PULL_REQUEST_TEMPLATE.md
    ├── workflows/
    │   ├── docs.yml
    │   ├── lint.yml
    │   └── security.yml
    ├── SECURITY.md
    ├── SUPPORT.md
    ├── FUNDING.yml
    ├── CODEOWNERS
    └── labels.yml
```

---
