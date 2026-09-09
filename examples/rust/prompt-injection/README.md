# Prompt Injection (CWE-1427)

Le code vulnérable concatène le prompt système et le message utilisateur dans une seule chaîne de texte brute envoyée au modèle, sans séparation structurelle des rôles, puis utilise la sortie directement sans validation. La correction transmet le prompt système et l'entrée utilisateur via les rôles structurés de l'API (`system`/`user`), et ajoute une validation applicative de la réponse avant tout usage, traitant la sortie du modèle comme une donnée non fiable. Élimine la classe de vulnérabilité CWE-1427 (Improper Neutralization of Input Used for LLM Prompting).
