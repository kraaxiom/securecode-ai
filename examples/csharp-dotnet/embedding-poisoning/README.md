# Embedding Poisoning (CWE-349)

La version vulnérable génère et indexe l'embedding de tout contenu soumis par un utilisateur (avis, description) dans la base vectorielle sans aucune modération préalable ni limite de fréquence, permettant à un même émetteur d'injecter en volume des vecteurs conçus pour polluer la recherche sémantique. La correction modère le contenu avant génération d'embedding, limite le nombre d'insertions par source sur une fenêtre de temps donnée, et journalise les insertions pour permettre une surveillance de la distribution des vecteurs indexés.
