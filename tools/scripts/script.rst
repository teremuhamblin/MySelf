=========================
Scripts — MySelf
=========================

Description
-----------
Ce dossier contient les scripts utilisés pour automatiser certaines actions du système MySelf.  
Chaque script suit une structure simple, claire et sécurisée.

Structure d’un script
---------------------
Tous les scripts utilisent un template standard :

- En-tête descriptif
- Sécurisation (`set -euo pipefail`)
- Fonctions utilitaires (`log`, `error`, `usage`)
- Fonction `main()` comme point d’entrée
- Exécution finale : ``main "$@"``

Objectifs
---------
- Automatiser des tâches répétitives
- Maintenir une structure cohérente
- Faciliter la maintenance et l’évolution

Template
--------
Un template réutilisable est fourni dans ``TEMPLATE.sh`` pour créer de nouveaux scripts.
