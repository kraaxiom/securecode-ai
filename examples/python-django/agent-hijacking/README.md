# Agent Hijacking — CWE-1427 / OWASP LLM01:2025 (Prompt Injection)

Le code vulnérable expose un agent LangChain qui planifie et exécute directement des outils à fort impact (`send_payment`, `delete_account`) via `AgentExecutor.invoke`, sans aucune validation métier indépendante ni confirmation humaine. Un contenu détourné (message utilisateur, ou document lu par un autre outil de l'agent) peut ainsi faire déclencher des actions irréversibles.

La correction sépare la couche de décision (le LLM produit un plan) de la couche d'exécution (l'application valide chaque étape). Les outils identifiés comme `HIGH_IMPACT_TOOLS` exigent une confirmation humaine explicite via `require_human_confirmation` avant exécution, et chaque appel d'outil passe par `validate_against_business_rules`, un contrôle indépendant du modèle. Toutes les décisions et actions sont journalisées via `audit_log` pour permettre l'audit post-incident.

Résiduel : cette correction ne dispense pas de restreindre également les permissions de l'agent lors du traitement de contenu externe non fiable (moindre privilège contextuel), ni de revoir périodiquement le catalogue `HIGH_IMPACT_TOOLS` à mesure que de nouveaux outils sont ajoutés. Voir `rules/remediation/agent-hijacking.md` pour d'autres langages.
