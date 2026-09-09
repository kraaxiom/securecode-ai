# secret-leakage (CWE-200)

## Description de la vulnérabilité

Ce pattern couvre les cas où des secrets (clés API, identifiants, jetons internes) se retrouvent exposés à travers l'usage d'un LLM, notamment lorsqu'ils sont inclus par erreur dans le prompt système transmis au modèle, ou lorsqu'un agent ayant accès au système de fichiers restitue le contenu de fichiers de configuration sensibles dans ses réponses.

Dans `vulnerable.go`, `buildSystemPrompt` insère la constante `internalAPIKey` (un placeholder, `REPLACE_WITH_YOUR_API_KEY`) directement en clair dans le texte du prompt système. `readConfigForAgent` donne par ailleurs à l'agent un accès libre à n'importe quel chemin du système de fichiers, sans restriction de scope. `getLLMResponse` journalise la conversation complète en clair et ne filtre jamais la sortie du modèle avant de la renvoyer.

## CWE réel utilisé

**CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor**, tel que documenté dans `knowledge/llm/secret-leakage.md` (OWASP LLM02:2025 – Sensitive Information Disclosure).

## Pourquoi c'est dangereux

- Un secret inclus dans un prompt système fait partie du contexte transmis au modèle et peut être restitué dans n'importe quelle réponse générée.
- Un accès fichier non restreint permet à l'agent de lire et potentiellement résumer des fichiers de configuration contenant d'autres secrets.
- L'absence de filtrage de sortie signifie qu'aucune barrière applicative ne détecte un secret avant qu'il n'atteigne l'utilisateur final.
- Les journaux de conversation en clair deviennent eux-mêmes une source de fuite persistante.

## Comment le correctif fonctionne

Le fichier `fixed.go` applique la remédiation décrite dans `rules/remediation/secret-leakage.md` :

1. **Référence indirecte** : `buildSystemPrompt` ne contient plus aucun secret ; l'agent est informé d'utiliser l'outil `callServiceX`, dont le secret réel est résolu côté application via `resolveSecret` (gestionnaire de secrets).
2. **Restriction d'accès fichier** : `readConfigForAgent` n'autorise que les répertoires listés dans `allowedConfigDirs`, excluant tout répertoire de configuration sensible.
3. **Filtrage de sortie** : `containsSecret` applique une détection par motif avant tout renvoi de réponse ; toute correspondance déclenche un blocage et une entrée d'audit.
4. **Journalisation sans secret** : les journaux ne consignent plus le contenu potentiellement sensible du prompt.

## Références

- OWASP Top 10 for LLM Applications : LLM02:2025 – Sensitive Information Disclosure
- CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor
- `rules/remediation/secret-leakage.md`
- `knowledge/llm/secret-leakage.md`
