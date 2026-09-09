## Secret Leakage via LLM (CWE-200)

Le code vulnérable inclut une clé API interne en clair directement dans le system prompt transmis au modèle et renvoie la réponse générée à l'utilisateur sans aucun filtrage, ce qui expose le secret à une restitution accidentelle si le modèle le reproduit dans sa réponse. La correction retire tout secret du prompt en le remplaçant par une référence indirecte à un outil dont le jeton est résolu uniquement côté application au moment de l'appel, et ajoute un filtrage de sortie qui détecte et bloque tout motif de secret avant de renvoyer la réponse au client.
