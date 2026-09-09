---
id: expression-language-injection
category: injections
cwe: CWE-917
owasp: A03:2021-Injection
severity_default: critical
languages: [java]
---

# Expression Language (EL) Injection

## Description
L'injection de langage d'expression cible les moteurs EL utilisés par les frameworks Java (JSF, Spring EL, OGNL, MVEL) pour évaluer dynamiquement des expressions dans les vues ou la configuration. Lorsqu'une entrée utilisateur est intégrée dans une expression évaluée par ce moteur, l'attaquant peut accéder à des objets Java arbitraires et, selon le contexte, exécuter du code sur le serveur.

## Où ça apparaît typiquement
- Champs de formulaire ou paramètres réinjectés dans une expression JSF/Spring EL/OGNL évaluée côté serveur.
- Struts et frameworks basés sur OGNL où des paramètres de requête influencent des expressions.
- Moteurs de règles métier ou de validation utilisant EL pour interpréter des conditions configurables par l'utilisateur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Construction dynamique d'une expression EL/OGNL/SpEL à partir d'une chaîne contenant une entrée utilisateur.
- Utilisation d'API d'évaluation d'expression (`ExpressionFactory`, `SpelExpressionParser`, `Ognl.getValue`) sur une chaîne non constante.
- Frameworks connus pour des CVE d'injection EL utilisés sans mise à jour de sécurité récente.

## Remédiation
- Ne jamais construire une expression EL à partir d'une entrée utilisateur ; utiliser des expressions statiques avec des paramètres liés séparément.
- Maintenir à jour les frameworks (Struts, Spring) pour bénéficier des correctifs contre les vecteurs d'injection EL connus.
- Restreindre le contexte d'évaluation (sandboxing, listes blanches de classes accessibles) quand l'évaluation dynamique est indispensable.
- Voir `rules/remediation/expression-language-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/expression-language-injection/`.

## Références
- OWASP Cheat Sheet: Injection Prevention Cheat Sheet
- CWE-917: Improper Neutralization of Special Elements used in an Expression Language Statement ('Expression Language Injection')
