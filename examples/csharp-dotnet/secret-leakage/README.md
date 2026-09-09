# Secret Leakage via LLM (CWE-200)

La version vulnérable construit dynamiquement le prompt système en y incluant en clair une clé API interne pour que l'agent puisse s'en servir, et renvoie/journalise les réponses du modèle sans aucun filtrage de sortie, ce qui peut exposer ce secret si le modèle le restitue dans sa réponse. La correction remplace le secret par une référence indirecte résolue côté application, et applique un filtrage de sortie qui détecte et masque tout motif de secret avant de renvoyer ou journaliser la réponse.
