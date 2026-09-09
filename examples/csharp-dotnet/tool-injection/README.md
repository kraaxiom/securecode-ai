# Tool Injection (CWE-1427)

La version vulnérable extrait le nom d'outil et les arguments directement du texte généré par le modèle puis les exécute tels quels, sans schéma de validation ni liste blanche contextuelle, ce qui permet à un contenu non fiable ayant influencé la génération de détourner l'appel d'outil. La correction valide chaque appel contre un schéma typé strict, restreint les outils disponibles via une liste blanche selon le niveau de confiance du contenu traité, et journalise chaque exécution.
