---
id: token-leakage
category: llm
cwe: CWE-522
owasp: LLM02:2025-Sensitive-Information-Disclosure
severity_default: high
languages: [python, js, java, csharp, go, php]
---

# Token Leakage (fuite de jetons d'authentification LLM)

## Description
Ce pattern concerne spécifiquement la fuite de jetons d'authentification liés aux services LLM eux-mêmes : clés d'API du fournisseur de modèle, jetons de session d'un assistant, ou identifiants OAuth utilisés par les plugins/outils connectés à un agent. Une exposition de ces jetons permet à un attaquant d'utiliser le service au nom de la victime (coût financier), d'accéder aux données associées au compte, ou de détourner les outils connectés à l'agent.

## Où ça apparaît typiquement
- Clés d'API de fournisseurs de modèles codées en dur côté client (application mobile, JavaScript frontend) au lieu d'être proxifiées côté serveur.
- Jetons OAuth de plugins/outils tiers stockés sans chiffrement ou journalisés en clair par erreur.
- Réponses d'erreur ou de débogage de l'intégration LLM incluant par inadvertance l'en-tête d'autorisation utilisé.
- Dépôts de code ou fichiers de configuration versionnés contenant des clés d'API de service LLM.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Clé d'API de fournisseur de modèle présente dans du code exécuté côté client (bundle JS, application mobile décompilable).
- Journalisation des en-têtes de requêtes complets sans masquage des jetons d'autorisation.
- Absence de rotation régulière des clés API et de scoping par environnement (dev/prod utilisant la même clé).
- Fichiers `.env` ou de configuration contenant des clés LLM présents dans l'historique d'un dépôt de code.

## Remédiation
- Ne jamais exposer de clé d'API de service LLM côté client ; toujours proxifier les appels via un backend qui détient le secret.
- Masquer systématiquement les jetons d'autorisation dans les journaux applicatifs et les traces de débogage.
- Mettre en place une rotation régulière des clés API avec scoping distinct par environnement et par service.
- Scanner les dépôts de code pour détecter la présence de secrets/jetons versionnés par erreur.
- Voir `rules/remediation/token-leakage.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/token-leakage/`.

## Références
- OWASP Top 10 for LLM Applications: LLM02:2025 – Sensitive Information Disclosure
- CWE-522: Insufficiently Protected Credentials
