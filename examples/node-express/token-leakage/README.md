## Token Leakage (CWE-522)

Le code vulnérable code en dur la clé API du fournisseur LLM et la renvoie au client via un endpoint de configuration, tout en journalisant les en-têtes de requête complets sans masquage, exposant le jeton à quiconque intercepte la réponse ou lit les journaux. La correction fait résoudre la clé depuis un gestionnaire de secrets côté serveur uniquement, proxifie tous les appels au modèle derrière un endpoint authentifié qui ne transmet jamais le secret au client, et applique un masquage systématique des en-têtes d'autorisation dans les journaux applicatifs.
