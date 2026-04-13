===============================
Générateur de Scripts — MySelf
===============================

Description
-----------
Le script ``generate.sh`` permet de créer automatiquement un nouveau script
basé sur le template standard ``TEMPLATE.sh``.

Usage
-----
Exécuter la commande suivante :

``./generate.sh <nom_du_script>``

Exemple :

``./generate.sh backup``

Cela crée un fichier :

``backup.sh``

Fonctionnement
--------------
Le générateur :

- copie le template
- remplace les champs ``<nom_du_script>`` et ``<description_courte>``
- crée un fichier exécutable prêt à l’emploi

Objectif
--------
Faciliter la création de scripts cohérents, propres et standardisés dans MySelf.
