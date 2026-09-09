# Token Leakage (CWE-522)

La version vulnérable inclut le jeton de session de l'utilisateur dans le contexte transmis au modèle pour personnaliser la réponse, et journalise l'en-tête d'autorisation complet à chaque appel, ce qui expose ces jetons dans les journaux applicatifs et le contexte du LLM. La correction exclut systématiquement les jetons d'authentification du contenu envoyé au modèle et masque les en-têtes d'autorisation dans les journaux et traces de télémétrie.
