# prompt-injection (CWE-1427)

## Description de la vulnérabilité

L'injection de prompt survient lorsqu'une entrée fournie par l'utilisateur parvient à modifier le comportement prévu d'un modèle de langage, faute de séparation structurelle entre les instructions de confiance (prompt système) et le contenu non fiable (entrée utilisateur). Ce répertoire illustre uniquement les faiblesses **architecturales** qui rendent cela possible — aucun texte d'attaque ou de contournement n'est inclus, conformément aux règles de ce skill défensif.

Dans `vulnerable.go`, la fonction `callLLM` construit le prompt par simple **concaténation de chaînes** (`systemPrompt + "\n" + userInput`) au lieu d'utiliser les rôles structurés (`system`/`user`) proposés par l'API du fournisseur de modèle. Le modèle ne dispose alors d'aucune frontière fiable entre l'instruction du développeur et le contenu utilisateur. De plus, `executeToolCalls` exécute directement tout appel d'outil renvoyé par le modèle, sans validation de schéma, sans distinction de niveau de privilège, et sans confirmation pour les actions sensibles (suppression de fichier, exécution de requêtes).

## CWE réel utilisé

**CWE-1427 : Improper Neutralization of Input Used for LLM Prompting**, tel que documenté dans `knowledge/llm/prompt-injection.md` (OWASP LLM01:2025 – Prompt Injection).

## Pourquoi c'est dangereux

- Sans séparation structurelle des rôles, le modèle ne peut pas distinguer de façon fiable une instruction légitime d'un contenu à traiter comme donnée.
- L'exécution automatique de tout appel d'outil renvoyé par le modèle, sans validation indépendante, transforme toute manipulation du comportement du modèle en action réelle sur le système (suppression, envoi, requête).
- L'absence de moindre privilège signifie que des outils à fort impact sont accessibles dans le même contexte que des outils bénins.

## Comment le correctif fonctionne

Le fichier `fixed.go` applique la remédiation décrite dans `rules/remediation/prompt-injection.md` :

1. **Rôles structurés de l'API** (`message{Role, Content}`) au lieu d'une concaténation de texte brut, pour que le fournisseur du modèle applique lui-même une frontière protocolaire entre instructions et contenu utilisateur.
2. **Validation de schéma indépendante** (`validateAgainstSchema`) : toute sortie du modèle est traitée comme une donnée non fiable et vérifiée avant toute exécution.
3. **Moindre privilège** : seuls des outils à faible impact (`lowImpactTools`) sont exposés par défaut ; les outils à fort impact (`highImpactTools`) exigent une confirmation humaine explicite (`requireHumanConfirmation`), refusée par défaut.
4. **Journalisation d'audit** (`auditLog`) de chaque tour de conversation et de chaque décision de blocage, pour permettre l'investigation après incident.

## Références

- OWASP Top 10 for LLM Applications : LLM01:2025 – Prompt Injection
- CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
- `rules/remediation/prompt-injection.md`
- `knowledge/llm/prompt-injection.md`
