# tool-injection (CWE-1427)

## Description de la vulnérabilité

L'injection d'outils (abus de function calling) survient lorsqu'un contenu non fiable parvient à influencer les paramètres ou la sélection des outils qu'un agent LLM invoque, ou lorsque la sortie d'un outil non fiable alimente les paramètres d'un outil suivant sans validation. Ce répertoire illustre uniquement la faiblesse **architecturale** en jeu : aucun appel d'outil malveillant réel n'est décrit.

Dans `vulnerable.go`, `executeToolCall` transmet directement les arguments générés par le modèle à l'implémentation réelle de l'outil (`toolRegistry`), sans validation de schéma, sans liste blanche contextuelle selon le niveau de confiance du contenu traité. `pipelineHandler` enchaîne plusieurs appels d'outils où le résultat de l'un (`lastResult`) est implicitement réinjecté dans le traitement suivant sans étape de nettoyage. `loadToolDefinitions` charge par ailleurs des définitions d'outils depuis une source externe arbitraire sans vérifier qu'elle est de confiance.

## CWE réel utilisé

**CWE-1427 : Improper Neutralization of Input Used for LLM Prompting**, tel que documenté dans `knowledge/llm/tool-injection.md` (OWASP LLM06:2025 – Excessive Agency).

## Pourquoi c'est dangereux

- Sans validation de schéma indépendante, tout paramètre généré par le modèle — potentiellement influencé par du contenu externe non fiable en amont du pipeline — est exécuté tel quel.
- Sans liste blanche contextuelle, tous les outils restent disponibles quel que soit le niveau de confiance du contenu en cours de traitement, y compris pour du contenu de faible confiance (sortie d'un outil précédent).
- Le chargement de définitions d'outils depuis des sources externes non vérifiées permet l'introduction de descriptions ou de métadonnées d'outils non fiables dans le registre applicatif.

## Comment le correctif fonctionne

Le fichier `fixed.go` applique la remédiation décrite dans `rules/remediation/tool-injection.md` :

1. **Validation de schéma stricte** (`validateArgs`) : chaque appel d'outil est vérifié contre un schéma typé (`toolSchemas`) avant exécution, indépendamment du texte généré par le modèle.
2. **Sources de confiance vérifiées** (`trustedToolSources`, `loadToolDefinitions`) : seules les définitions d'outils provenant d'un registre interne vérifié sont acceptées.
3. **Liste blanche contextuelle** (`allowedToolsByTrust`) : le niveau de confiance du contenu traité (`trustUserDirect` vs `trustLowConfidence`) restreint dynamiquement l'ensemble des outils exécutables — dans `pipelineHandler`, la confiance est explicitement abaissée dès qu'un résultat d'outil précédent influence l'appel suivant.
4. **Journalisation d'audit** (`auditLog`) de chaque appel d'outil réellement exécuté et de chaque rejet, pour permettre la détection d'anomalies après incident.

## Références

- OWASP Top 10 for LLM Applications : LLM06:2025 – Excessive Agency
- CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
- `rules/remediation/tool-injection.md`
- `knowledge/llm/tool-injection.md`
