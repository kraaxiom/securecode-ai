---
id: open-redirect
category: frontend
cwe: CWE-601
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [js, ts, php, python, java]
---

# Redirection ouverte (Open Redirect)

## Description
Une redirection ouverte survient quand une application redirige l'utilisateur vers une URL fournie en paramètre sans en valider la destination, permettant à un attaquant de construire un lien vers le domaine de confiance qui redirige ensuite vers un site malveillant. Cette technique est largement utilisée dans les campagnes de phishing car le lien initial pointe vers un domaine légitime, contournant la méfiance de l'utilisateur et certains filtres anti-phishing.

## Où ça apparaît typiquement
- Paramètres de redirection post-authentification (`?redirect=`, `?next=`, `?returnUrl=`, `?continue=`).
- Liens de déconnexion redirigeant vers une URL de retour fournie par le client.
- Passerelles de tracking/liens sortants qui redirigent via une URL en paramètre de requête.
- Intégrations SSO ou OAuth où l'URL de callback n'est pas strictement validée contre une liste blanche.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction de redirection (`res.redirect(...)`, `header('Location: ...')`, `window.location = ...`) avec une valeur directement issue d'un paramètre de requête.
- Validation de destination basée uniquement sur un `startsWith('/')` insuffisant (contournable par des URL protocol-relative `//evil.com`).
- Absence de liste blanche de domaines ou de chemins autorisés pour les redirections post-action.
- Comparaison de domaine par sous-chaîne (`includes('monsite.com')`) plutôt que par égalité stricte de l'hôte.

## Remédiation
- Valider les URL de redirection contre une liste blanche stricte de chemins relatifs ou de domaines autorisés.
- Refuser les URL absolues ou protocol-relative fournies par l'utilisateur pour les redirections internes.
- Préférer des identifiants indirects (ex: un code mappé côté serveur vers une URL) plutôt que l'URL complète en paramètre.
- Voir `rules/remediation/open-redirect.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/open-redirect/`.

## Références
- OWASP Cheat Sheet: Unvalidated Redirects and Forwards Cheat Sheet
- CWE-601: URL Redirection to Untrusted Site ('Open Redirect')
