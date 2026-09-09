# Agent Hijacking (détournement d'agent)

## Description de la vulnérabilité

Le détournement d'agent désigne le scénario où un agent LLM autonome, disposant d'un accès à des outils réels (paiement, suppression de compte, envoi d'e-mail, exécution de commandes), exécute directement les actions décidées par le modèle sans aucune validation applicative intermédiaire. Un attaquant qui parvient à influencer la décision du modèle — via une injection de prompt directe ou indirecte — peut alors faire déclencher des actions non prévues dans des systèmes externes réels.

Dans `vulnerable.go`, la fonction `RunAgent` demande un plan au LLM puis appelle `executeTool` pour chaque étape retournée, sans distinction entre outils anodins et outils à fort impact, sans confirmation humaine, et sans journalisation.

## CWE réel utilisé

**CWE-1427 : Improper Neutralization of Input Used for LLM Prompting** (tel que référencé dans `knowledge/llm/agent-hijacking.md`).

## Pourquoi c'est dangereux

- La gravité dépasse une simple réponse textuelle incorrecte : l'agent a une capacité d'action réelle (paiement, suppression, envoi externe).
- Aucune séparation entre la couche de décision (LLM) et la couche d'exécution (application) : le modèle est de facto l'unique autorité pour déclencher des opérations irréversibles.
- L'agent conserve les mêmes privilèges élevés même lorsqu'il traite du contenu externe non fiable, ce qui ouvre la voie à un détournement via injection de prompt indirecte.
- L'absence de journalisation empêche toute investigation ou détection d'anomalie après incident.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Une liste explicite `highImpactTools` regroupant les outils sensibles (paiement, suppression de compte).
2. Une confirmation humaine obligatoire (`HumanConfirmer.Confirm`) avant toute exécution d'une action à fort impact.
3. Une validation métier indépendante du modèle (`BusinessValidator.Validate`) appliquée à chaque étape, quel que soit l'outil.
4. Un cloisonnement contextuel : lorsque `UntrustedContext` est vrai (l'agent traite du contenu externe non fiable), les outils à fort impact sont bloqués d'office, réduisant les privilèges de l'agent dans ce contexte.
5. Une journalisation systématique (`AuditLogger.Record`) de chaque décision, blocage et exécution, permettant l'audit post-incident.

## Références

- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
- `rules/remediation/agent-hijacking.md`
- `knowledge/llm/agent-hijacking.md`
