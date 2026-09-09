# token-leakage (CWE-522)

## Description de la vulnérabilité

Ce pattern concerne la fuite de jetons d'authentification liés aux services LLM eux-mêmes (clés d'API du fournisseur de modèle). Une exposition de ces jetons permet à un attaquant d'utiliser le service au nom de la victime (coût financier), d'accéder aux données associées au compte, ou de détourner les outils connectés à l'agent.

Dans `vulnerable.go`, la constante `llmProviderAPIKey` (placeholder `REPLACE_WITH_YOUR_API_KEY`) est codée en dur dans le code backend. `debugHandler` journalise en clair l'intégralité des en-têtes de réponse, y compris tout en-tête d'autorisation, sans aucun masquage — un pattern qui, en production, expose le jeton dans les journaux applicatifs et les traces de débogage.

## CWE réel utilisé

**CWE-522 : Insufficiently Protected Credentials**, tel que documenté dans `knowledge/llm/token-leakage.md` (OWASP LLM02:2025 – Sensitive Information Disclosure).

## Pourquoi c'est dangereux

- Une clé codée en dur dans le code source finit presque systématiquement par se retrouver dans l'historique d'un dépôt de code ou un artefact de build.
- Une journalisation de débogage qui inclut les en-têtes complets expose l'en-tête `Authorization` en clair dans des journaux souvent moins protégés que le code source lui-même.
- L'absence de rotation et de scoping par environnement signifie qu'une seule fuite compromet l'ensemble des environnements (dev, staging, prod).

## Comment le correctif fonctionne

Le fichier `fixed.go` applique la remédiation décrite dans `rules/remediation/token-leakage.md` :

1. **Résolution via gestionnaire de secrets** : `resolveSecret` lit la clé depuis une variable d'environnement injectée par un secret manager, jamais codée en dur (le repli local reste un placeholder explicite).
2. **Masquage systématique des journaux** : `redactHeaders` remplace toute valeur d'en-tête sensible (`Authorization`, `X-Api-Key`) par `[REDACTED]` avant toute journalisation.
3. **Architecture en proxy backend** : seul le backend (`askAssistant`) détient la clé du fournisseur LLM ; aucun client ne la reçoit jamais, et chaque appel au backend est lui-même authentifié (`requireAuth`) indépendamment du secret du fournisseur.

## Références

- OWASP Top 10 for LLM Applications : LLM02:2025 – Sensitive Information Disclosure
- CWE-522 : Insufficiently Protected Credentials
- `rules/remediation/token-leakage.md`
- `knowledge/llm/token-leakage.md`
